<template>
  <div>
    <div class="p-fluid">
      <div class="p-field field-margin">
        <span class="p-float-label p-my-3">
          <Password v-model="password" id="password" toggleMask :feedback="false"></Password>
          <label for="password">Enter Secret Password</label>
        </span>
      </div>
      <div class="p-field field-margin">
        <span class="p-float-label p-my-3">
          <Dropdown v-model="ttlAmountDays" id="ttl" :options="ttlOptions" optionLabel="label" optionValue="value" placeholder="items"></Dropdown>
          <label for="ttl">Link will expire after</label>
        </span>
      </div>
      <div class="p-field">
        <span class="p-float-label p-my-3">
          <Dropdown v-model="maxRevielCount" id="maxReviel" :options="maxRevielOptions" optionLabel="label" optionValue="value" placeholder="items"></Dropdown>
          <label for="ttl">Link can be used</label>
        </span>
      </div>
      <Button v-on:click="generateLink" :disabled="password === '' || (maxRevielCount === -1 && ttlAmountDays === -1)">Generate Link</Button> 
      <div v-if="maxRevielCount === -1 && ttlAmountDays === -1" class="btn-err">It is required to specify at least one exiration criteria.</div>
      <div>
        <p>
          An <strong>encrypted link</strong> will be generated.
        </p> 
        <p>
          It can be used to <strong>share passwords</strong> via E-Mail securely. 
        </p>
        <!-- <p v-if="maxRevielCount > 0" >
          <span v-if="maxRevielCount === 1">The link can be used <strong>only once!</strong></span>
          <span v-if="maxRevielCount > 1">The link can be used <strong>only {{ maxRevielCount }} times!</strong></span>
        </p>
        <p v-if="ttlAmountDays > 0" >
          <span v-if="ttlAmountDays === 1">It expires after <strong>one Day!</strong></span>
          <span v-if="ttlAmountDays > 1">It expires after <strong>only {{ ttlAmountDays }} Days!</strong></span>
        </p> -->
        <p>
          The secret is <strong>never transmited</strong> to the server in clear text.
        </p>
      </div>


      <span class="p-inputgroup" v-if="generatedLink !== ''">
        <InputText type="text" v-model="generatedLink" disabled/>
        <span class="p-inputgroup-addon" v-on:click="copyLink">
          <i class="pi pi-copy"></i>
        </span>
      </span>
      <div v-if="copied" class="p-pt-3">
        <InlineMessage severity="success">Copied to clipboard</InlineMessage>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { Service, OpenAPI } from '../api'
import { Encryption } from '../crypto/encryption'
import { EncodingUrlSafe } from '../crypto/encoding-url-safe'
import copy from 'copy-to-clipboard'

export default defineComponent({
  name: 'EnterSecret',
  data () {
    return {
      password: '',
      ttlAmountDays: 30,
      ttlOptions: [
        { value: 1, label: "24 Hours" },
        { value: 7, label: "7 Days" },
        { value: 30, label: "30 Days" },
        { value: 90, label: "90 Days" },
        { value: 365, label: "365 Days" },
        { value: -1, label: "Unlimited" },
      ],
      ttlUnit: 'Days',
      maxRevielCount: 1,
      maxRevielOptions: [
        { value: 1, label: "Only once" },
        { value: 3, label: "3 Times" },
        { value: 10, label: "10 Times" },
        { value: -1, label: "Unlimited" },
      ],
      generatedLink: '',
      copied: false,
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
        timeToLive: this.ttlAmountDays > 0 ? this.ttlAmountDays * 24 * 60 * 60 : -1,
        maxRetrievalCount: this.maxRevielCount,
      })

      const keyBytes = await Encryption.exportKey(key)
      const keyString = EncodingUrlSafe.encode(keyBytes)

      this.generatedLink = window.location.protocol + '//' + window.location.host + '/sec/' + resp.id + '#' + keyString

      } catch (ex) {
        console.log( "ex ", ex)
      }
    },
    copyLink() {
      if (this.generatedLink && this.generatedLink !== '') {
        this.copied = copy(this.generatedLink)
        console.log("copy success: ", this.copied)
      }
    }
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.field-margin {
  margin-bottom: 2rem;
}
.btn-err {
  color: red;
  margin-top: 5px;
  margin-left: 5px;
}

</style>
