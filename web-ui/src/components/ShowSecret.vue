<template>
  <div> 
    <div v-if="!dataLoaded">
      Access to the secret expired!<br>
      <a href="/">Enter a new secret</a>
    </div>
    <div v-if="dataLoaded">
      This secret will self destruct after <strong>{{ expirationDate.toLocaleDateString() }} {{ expirationDate.toLocaleTimeString() }}</strong><br>
      or after showing it <strong>{{ maxRevielCount - revielCount }}</strong> more times (<strong>{{ revielCount }}/{{ maxRevielCount }}</strong>).<br> <br>
      <button v-on:click="showSecret" :disabled="secretShown">Show Secret</button><br> <br>
      <strong v-if="secretShown && secret !== ''">{{ secret }}</strong>
      <strong v-if="secretShown && secret === ''">Your link appears to be broken.</strong>
    </div>

  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Service, OpenAPI } from '../api'
import { EncodingUrlSafe } from '../crypto/encoding-url-safe'
import { Encryption, IEncryptedPayload } from '../crypto/encryption'

export default defineComponent({
  name: 'ShowSecret',
  async created() {
    try {
      OpenAPI.BASE = '/api/v1'
      this.id = this.$route.params.id as string
      const hash = this.$route.hash as string
      if (hash.startsWith('#')) {
        this.key = hash.slice(1)
      }

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
      key: "",
    }
  },
  methods: {
    async showSecret () {
      try {
        let secret = await Service.revealSecret(this.id)

        const encrypted: IEncryptedPayload = JSON.parse(atob(secret.encryptedSecret ?? ""))
        const keyBytes = EncodingUrlSafe.decode(this.key)
        const key = await Encryption.importKey(keyBytes)
        const plaintext = await Encryption.decrypt(encrypted, key)

        this.secret = plaintext
        this.revielCount++
        this.secretShown = true
      } catch(ex) {
        console.error("Failed to decrypt")
      }

    }
  },
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
