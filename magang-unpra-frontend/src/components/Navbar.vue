<template>
  <nav class="fixed top-0 left-0 right-0 w-full z-50 shadow-md">
    <!-- Baris 1: Logo -->
    <div class="w-full bg-white px-8 py-3 flex items-center">
      <RouterLink to="/" class="flex items-center gap-4">
        <img src="/images/logo-telpp.png" alt="TeL Logo" class="h-12 w-auto" />
        <span class="text-sm font-bold text-gray-800 tracking-widest uppercase">
          PT Tanjungenim Lestari Pulp and Paper
        </span>
      </RouterLink>
    </div>

    <!-- Baris 2: Menu hijau -->
    <div class="w-full bg-[#5F9E42] flex items-center px-8">

      <!-- Our Company -->
      <RouterLink to="/our-company"
        class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap"
        :class="{ 'bg-white/30 border-b-2 border-white': isActive('/our-company') }">
        OUR COMPANY
      </RouterLink>

      <!-- Product -->
      <RouterLink to="/product"
        class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap"
        :class="{ 'bg-white/30 border-b-2 border-white': isActive('/product') }">
        PRODUCT
      </RouterLink>

      <!-- News -->
      <RouterLink to="/news"
        class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap"
        :class="{ 'bg-white/30 border-b-2 border-white': isActive('/news') }">
        NEWS
      </RouterLink>

      <!-- Sustainability -->
      <div class="relative group">
        <button
          class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 flex items-center gap-1.5 hover:bg-white/20 transition whitespace-nowrap"
          :class="{ 'bg-white/30 border-b-2 border-white': isActive('/sustainability') }">
          SUSTAINABILITY <span class="text-[9px]">▾</span>
        </button>
        <div class="absolute top-full left-0 bg-white border border-gray-200 min-w-[230px] z-50 hidden group-hover:block shadow-lg">
          <RouterLink to="/sustainability/forest-management"
            class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100"
            :class="{ 'bg-green-50 text-[#5F9E42] font-bold': isActive('/sustainability/forest-management') }">
            Forest Management
          </RouterLink>
          <RouterLink to="/sustainability/people-development"
            class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100"
            :class="{ 'bg-green-50 text-[#5F9E42] font-bold': isActive('/sustainability/people-development') }">
            People Development
          </RouterLink>
          <RouterLink to="/sustainability/supply-chain"
            class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100"
            :class="{ 'bg-green-50 text-[#5F9E42] font-bold': isActive('/sustainability/supply-chain') }">
            Supply Chain
          </RouterLink>
          <RouterLink to="/sustainability/pulp-process"
            class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100"
            :class="{ 'bg-green-50 text-[#5F9E42] font-bold': isActive('/sustainability/pulp-process') }">
            Pulp Process
          </RouterLink>
          <RouterLink to="/sustainability/safety-health"
            class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100"
            :class="{ 'bg-green-50 text-[#5F9E42] font-bold': isActive('/sustainability/safety-health') }">
            Safety &amp; Health
          </RouterLink>
          <div class="relative group/csr">
            <button
              class="w-full text-left px-5 py-3 text-xs hover:bg-green-50 hover:text-[#5F9E42] flex justify-between items-center border-b border-gray-100"
              :class="isActive('/sustainability/csr') ? 'bg-green-50 text-[#5F9E42] font-bold' : 'text-gray-700'">
              Corporate Social Responsibility <span>›</span>
            </button>
            <div class="absolute left-full top-0 bg-white border border-gray-200 min-w-[210px] z-50 hidden group-hover/csr:block shadow-lg">
              <RouterLink to="/sustainability/csr/vision"
                class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100"
                :class="{ 'bg-green-50 text-[#5F9E42] font-bold': isActive('/sustainability/csr/vision') }">
                Vision And Mission
              </RouterLink>
              <RouterLink to="/sustainability/csr/community"
                class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100"
                :class="{ 'bg-green-50 text-[#5F9E42] font-bold': isActive('/sustainability/csr/community') }">
                Local Community Development
              </RouterLink>
              <RouterLink to="/sustainability/csr/report"
                class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42]"
                :class="{ 'bg-green-50 text-[#5F9E42] font-bold': isActive('/sustainability/csr/report') }">
                CSR Report
              </RouterLink>
            </div>
          </div>
        </div>
      </div>

      <!-- Biodiversity -->
      <a href="https://kehati.telpp.com/" target="_blank"
        class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap">
        BIODIVERSITY
      </a>

      <!-- Contact -->
    <RouterLink to="/contact" class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap">
     CONTACT
    </RouterLink>

      <div class="flex-1"></div>
    </div>
  </nav>
</template>

<script setup>
import { RouterLink, useRoute } from 'vue-router'

const route = useRoute()

const isActive = (path) => {
  if (path === '/') return route.path === '/'
  return route.path.startsWith(path)
}
</script>