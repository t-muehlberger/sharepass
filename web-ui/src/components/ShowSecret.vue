<template>
  <div class="p-fluid">  
    <div v-if="!dataLoaded">
      Access to the secret expired!
      <br>
      <br>
      <a href="/">Enter a new secret</a>
    </div>
    <div v-if="dataLoaded">
      <div v-if="!expirationDisabled"> This secret will self destruct after <strong>{{ expirationDate.toLocaleDateString() }} {{ expirationDate.toLocaleTimeString() }}</strong><br></div>
      <div v-if="!expirationDisabled && !allowUnlimitedReviel">or after showing it <strong>{{ (maxRevielCount || 0) - (revielCount || 0) }}</strong> more times (<strong>{{ revielCount }}/{{ maxRevielCount }}</strong>).</div> 
      <div v-if="expirationDisabled && !allowUnlimitedReviel">This secret will self destruct after showing it <strong>{{ (maxRevielCount || 0) - (revielCount || 0) }}</strong> more times (<strong>{{ revielCount }}/{{ maxRevielCount }}</strong>).</div>
      <br>
      <Button v-on:click="showSecret" :disabled="secretShown">Show Secret</Button><br> <br>
      <strong v-if="secretShown && secret === ''">Your link appears to be broken.</strong>

      <span class="p-inputgroup">
        <InputText type="text" v-model="secret" disabled/>
        <span class="p-inputgroup-addon" v-on:click="copySecret()">
          <i class="pi pi-copy"></i>
        </span>
      </span>
    </div>
    <div v-if="copied" class="p-pt-3">
      <InlineMessage severity="success" >Copied to clipboard</InlineMessage>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Service, OpenAPI } from '../api'
import { EncodingUrlSafe } from '../crypto/encoding-url-safe'
import { Encryption, IEncryptedPayload } from '../crypto/encryption'
import copy from 'copy-to-clipboard'

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
      this.allowUnlimitedReviel = secretMetadata.maxRetrievalCount === undefined || secretMetadata.maxRetrievalCount === null
      this.expirationDisabled = secretMetadata.expiryTime === undefined || secretMetadata.expiryTime === undefined
      this.maxRevielCount = secretMetadata.maxRetrievalCount || 0
      this.revielCount = secretMetadata.retrievalCount || 0
      this.dataLoaded = true
    } catch { }
  },
  data () {
    return {
      id: "",
      dataLoaded: false,
      secretShown: false,
      expirationDisabled: false,
      expirationDate: new Date(),
      allowUnlimitedReviel: false,
      maxRevielCount: 0 as number,
      revielCount: 0 as number,
      secret: "",
      key: "",
      copied: false,
    }
  },
  methods: {
    async showSecret () {
      try {
        let secret = await Service.revealSecret(this.id)

        const encrypted: IEncryptedPayload = {
          encrypted: secret.encryptedSecret ?? "",
          initializationVector: secret.initializationVector ?? "",
        }
        const keyBytes = EncodingUrlSafe.decode(this.key)
        const key = await Encryption.importKey(keyBytes)
        const plaintext = await Encryption.decrypt(encrypted, key)

        this.secret = plaintext
        if (typeof(this.revielCount) === "number") {
          this.revielCount++
        }
        this.secretShown = true
      } catch(ex) {
        console.error("Failed to decrypt")
      }
    },
    copySecret() {
      if (this.secret && this.secret !== '') {
        this.copied = copy(this.secret)
        console.log("Copy success ", this.copied)
      }
    }
  },
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
