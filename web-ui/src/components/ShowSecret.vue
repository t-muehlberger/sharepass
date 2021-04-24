<template>
  <div> 
    <div v-if="!dataLoaded">
      Access to the secret expired!<br>
      <a href="/">Enter a new secret</a>
    </div>
    <div v-if="dataLoaded">
      This secret will self destruct after <strong>{{ expirationDate.toLocaleDateString() }} {{ expirationDate.toLocaleTimeString() }}</strong><br>
      or after showing it <strong>{{ maxRevielCount - revielCount }}</strong> more times (<strong>{{ revielCount }}/{{ maxRevielCount }}</strong>).<br>
      <button v-on:click="showSecret" :disabled="secretShown">Show Secret</button><br>
      <strong>{{ secret }}</strong>
    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Service, OpenAPI } from '../api'
import { Crypter } from '../crypto/crypter'

export default defineComponent({
  name: 'ShowSecret',
  async created() {
    try {
      OpenAPI.BASE = '/api/v1'
      this.id = this.$route.params.id as string

      let secretMetadata = await Service.getSecretMetadata(this.id)
      this.expirationDate = new Date(secretMetadata.expiryTime ?? "")
      this.maxRevielCount = secretMetadata.maxRetrievalCount ?? 0
      this.revielCount = secretMetadata.retrievalCount ?? 0
      this.dataLoaded = true
    } catch { }
  },
  data () {
    return {
      id: "",
      dataLoaded: false,
      secretShown: false,
      expirationDate: new Date(),
      maxRevielCount: 0,
      revielCount: 0,
      secret: "",
    }
  },
  methods: {
    async showSecret () {
      let secret = await Service.revealSecret(this.id)

      const ciphertext = atob(secret.encryptedSecret ?? "")
      const key = "supersecret"
      const plaintext = Crypter.decrypt(ciphertext, key)

      this.secret = plaintext
      this.revielCount++
      this.secretShown = true
    }
  },
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
