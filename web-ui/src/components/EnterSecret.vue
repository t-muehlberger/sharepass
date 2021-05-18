<template>
  <div class="">
    <label for="in-pwd">Enter Secret</label><br>
    <input v-model="password" placeholder="Secret" type="password" id="input-pwd" name="input-pwd">
    <p>
      Secret will be available for the next <strong>{{ ttlAmount }} {{ ttlUnit }}</strong>.<br>
      It can be viewed a maximum of <strong>{{ maxRevielCount }}</strong> times.
    </p>
    <button v-on:click="generateLink" :disabled="password === ''">Generate Link</button> <br> <br>
    <a v-if="generatedLink !== ''" v-bind:href="generatedLink" target="_blank">{{ generatedLink }}</a>
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
