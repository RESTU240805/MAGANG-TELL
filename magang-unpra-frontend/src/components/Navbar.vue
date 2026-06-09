<template>
  <nav class="fixed top-0 w-full z-50 bg-white shadow-md">
    <div class="max-w-7xl mx-auto px-6 py-3 flex items-center justify-between">
      <RouterLink to="/" class="text-2xl font-black text-green-600">TeLpp.</RouterLink>

      <div class="hidden md:flex items-center gap-6 text-sm font-medium text-gray-700">
        <RouterLink to="/" class="hover:text-green-600 transition">Our Company</RouterLink>
        <RouterLink to="/product" class="hover:text-green-600 transition">Product</RouterLink>
        <RouterLink to="/news" class="hover:text-green-600 transition">News</RouterLink>

        <!-- Sustainability Dropdown -->
        <div class="relative" @mouseenter="showDropdown = true" @mouseleave="showDropdown = false">
          <button class="flex items-center gap-1 hover:text-green-600 transition">
            Sustainability
            <span class="text-xs">▾</span>
          </button>
          <div v-if="showDropdown"
            class="absolute top-full left-0 mt-2 w-56 bg-white rounded-xl shadow-lg border border-gray-100 py-2 z-50">
            <RouterLink to="/sustainability/forest-management"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600">
              Forest Management
            </RouterLink>
            <RouterLink to="/sustainability/people-development"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600">
              People Development
            </RouterLink>
            <RouterLink to="/sustainability/supply-chain"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600">
              Supply Chain
            </RouterLink>
            <RouterLink to="/sustainability/pulp-process"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600">
              Pulp Process
            </RouterLink>
            <RouterLink to="/sustainability/safety-health"
              class="block px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600">
              Safety & Health
            </RouterLink>
            <!-- CSR Sub-dropdown -->
            <div class="relative" @mouseenter="showCSR = true" @mouseleave="showCSR = false">
              <button class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600 flex justify-between items-center">
                Corporate Social Responsibility <span>›</span>
              </button>
              <div v-if="showCSR"
                class="absolute left-full top-0 w-52 bg-white rounded-xl shadow-lg border border-gray-100 py-2">
                <RouterLink to="/sustainability/csr/vision"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600">
                  Vision And Mission
                </RouterLink>
                <RouterLink to="/sustainability/csr/community"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600">
                  Local Community Development
                </RouterLink>
                <RouterLink to="/sustainability/csr/report"
                  class="block px-4 py-2 text-sm text-gray-700 hover:bg-green-50 hover:text-green-600">
                  CSR Report
                </RouterLink>
              </div>
            </div>
          </div>
        </div>

        <a href="#" class="hover:text-green-600 transition">Biodiversity</a>
        <a href="#" class="border border-gray-400 px-4 py-1.5 rounded hover:bg-gray-100 transition">E-Procurement</a>
        <a href="#" class="bg-green-600 text-white px-4 py-1.5 rounded hover:bg-green-700 transition">E-Recruitment</a>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref } from 'vue'
import { RouterLink } from 'vue-router'

const showDropdown = ref(false)
const showCSR = ref(false)
</script>