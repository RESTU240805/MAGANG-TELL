<template>
  <div class="min-h-screen">
    <Navbar v-if="!isAdminRoute" />
    <div :class="!isAdminRoute ? 'pt-[112px]' : ''">
      <RouterView />
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'

const route = useRoute()

const isAdminRoute = computed(() => route.path.startsWith('/admin'))
</script>