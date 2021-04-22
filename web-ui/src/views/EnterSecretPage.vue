<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <EnterSecret/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import EnterSecret from '@/components/EnterSecret.vue' // @ is an alias to /src

export default defineComponent({
  name: 'Home',
  components: {
    EnterSecret
  }
})
</script>
