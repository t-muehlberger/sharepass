<template>
  <div class="home">
    <EnterSecret/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import EnterSecret from '../components/EnterSecret.vue'

export default defineComponent({
  name: 'Home',
  components: {
    EnterSecret
  }
})
</script>
