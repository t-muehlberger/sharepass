<template>
  <div class="">
    <label for="in-pwd">Enter Secret</label><br>
    <input v-model="password" placeholder="Secret" type="password" id="input-pwd" name="input-pwd">
    <p>
      Secret will be available for the next <strong>{{ ttlAmount }} {{ ttlUnit }}</strong>.<br>
      It can be viewed a maximum of <strong>{{ maxRevielCount }}</strong> times.
    </p>
    <button v-on:click="generateLink" :disabled="password === ''">Generate Link</button> <br>
    <a v-if="generatedLink !== ''" v-bind:href="generatedLink" target="_blank">{{ generatedLink }}</a>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Service, OpenAPI } from '../api'
import { Crypter } from '../crypto/crypter'

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

      const key = Crypter.randomKey()
      const ecnrypted = Crypter.encrypt(this.password, key)

      const resp = await Service.createSecret({
        encryptedSecret: btoa(ecnrypted),
        timeToLive: this.ttlAmount * 24 * 60 * 60,
        maxRetrievalCount: this.maxRevielCount,
      })
      this.generatedLink = 'http://localhost:8080/sec/' + resp.id + '#' + key
    }
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
