<template>
  <div class="min-h-screen">
    <Navbar v-if="!isAdminRoute" />
    <RouterView />
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRoute, RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'

const route = useRoute()

// Sembunyikan navbar di semua halaman /admin/*
const isAdminRoute = computed(() => route.path.startsWith('/admin'))
</script>