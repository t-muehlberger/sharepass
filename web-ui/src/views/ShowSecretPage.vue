<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <ShowSecret/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import ShowSecret from '@/components/ShowSecret.vue'

export default defineComponent({
  name: 'Home',
  components: {
    ShowSecret
  }
})
</script>
