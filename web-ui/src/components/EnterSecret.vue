<template>
  <div class="card">
    <div class="p-fluid">
      <div class="p-field">
        <span class="p-float-label">
          <Password v-model="password" id="password" toggleMask :feedback="false"></Password>
          <label for="password">Enter Secret</label>
        </span>
      </div>
      <p>
        Secret will be available for the next <strong>{{ ttlAmount }} {{ ttlUnit }}</strong>.<br>
        It can be viewed a maximum of <strong>{{ maxRevielCount }}</strong> times.
      </p>
      <Button v-on:click="generateLink" :disabled="password === ''">Generate Link</Button> <br> <br>

      <span class="p-inputgroup" v-if="generatedLink !== ''">
        <InputText type="text" v-model="generatedLink" disabled/>
        <span class="p-inputgroup-addon">
          <i class="pi pi-copy"></i>
        </span>
      </span>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Service, OpenAPI } from '../api'
import { Encryption } from '../crypto/encryption'
import { EncodingUrlSafe } from '../crypto/encoding-url-safe'

export default defineComponent({
  name: 'EnterSecret',
  data () {
    return {
      password: '',
      ttlAmount: 7,
      ttlUnit: 'Days',
      maxRevielCount: 3,
      generatedLink: ''
    }
  },
  methods: {
    async generateLink () {
      OpenAPI.BASE = '/api/v1'

      try {

      const key = await Encryption.randomKey()
      const ecnrypted = await Encryption.encrypt(this.password, key)

      const resp = await Service.createSecret({
        encryptedSecret: ecnrypted.encrypted,
        initializationVector: ecnrypted.initializationVector,
        timeToLive: this.ttlAmount * 24 * 60 * 60,
        maxRetrievalCount: this.maxRevielCount,
      })

      const keyBytes = await Encryption.exportKey(key)
      const keyString = EncodingUrlSafe.encode(keyBytes)

      this.generatedLink = window.location.protocol + '//' + window.location.host + '/sec/' + resp.id + '#' + keyString

      } catch (ex) {
        console.log( "ex ", ex)
      }
    }
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
