<template>
  <div>
    <Menubar :model="items">
      <template #start>
        <strong>Sharepass</strong>
        <span> - Secure password sharing</span>
      </template>
    </Menubar>
    <router-view/>
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
