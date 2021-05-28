<template>
<div class="p-grid p-jc-center p-p-3">
  <div class="p-col-12 p-md-5 p-lg-5 p-xl-4">
    <Card>
      <template #title>
        Sharepass
      </template>
      <template #subtitle>
        Secure password sharing
      </template>
      <template #content>
        <router-view/>
      </template>
    </Card>
  </div>
</div>
</template>
<script lang="ts">
import { defineComponent } from "@vue/runtime-core"

export default defineComponent({
  name: 'EnterSecret',
  data () {
    return {
      items: [
        {
          label: 'Sharepass',
          visible: false
        },
      ]
    }
  },
  methods: {
  }
})
</script>
<style lang="scss">
</style>
