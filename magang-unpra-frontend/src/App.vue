<template>
  <div class="min-h-screen">
    <Navbar v-if="!isAdminRoute" />
    <div :class="!isAdminRoute ? 'pt-[112px]' : ''">
      <RouterView v-slot="{ Component }">
        <Transition name="page">
          <component :is="Component" />
        </Transition>
      </RouterView>
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

<style>
.page-enter-active {
  transition: opacity 0.25s ease;
}

.page-leave-active {
  transition: opacity 0.15s ease;
  position: absolute;
  width: 100%;
}

.page-enter-from {
  opacity: 0;
}

.page-leave-to {
  opacity: 0;
}
</style>