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

    <!-- Baris 2: Menu hijau full width -->
    <div class="w-full bg-[#5F9E42] flex items-center px-8">

      <!-- Our Company -->
      <RouterLink to="/" class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap">
        OUR COMPANY
      </RouterLink>

      <!-- Product -->
    <RouterLink to="/product" class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap">
      PRODUCT
    </RouterLink>

      <!-- News -->
      <RouterLink to="/news" class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap">
        NEWS
      </RouterLink>

      <!-- Sustainability -->
      <div class="relative group">
        <button class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 flex items-center gap-1.5 hover:bg-white/20 transition whitespace-nowrap">
          SUSTAINABILITY <span class="text-[9px]">▾</span>
        </button>
        <div class="absolute top-full left-0 bg-white border border-gray-200 min-w-[230px] z-50 hidden group-hover:block">
          <RouterLink to="/sustainability/forest-management" class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100">Forest Management</RouterLink>
          <RouterLink to="/sustainability/people-development" class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100">People Development</RouterLink>
          <RouterLink to="/sustainability/supply-chain" class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100">Supply Chain</RouterLink>
          <RouterLink to="/sustainability/pulp-process" class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100">Pulp Process</RouterLink>
          <RouterLink to="/sustainability/safety-health" class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100">Safety &amp; Health</RouterLink>
          <div class="relative group/csr">
            <button class="w-full text-left px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] flex justify-between items-center">
              Corporate Social Responsibility <span>›</span>
            </button>
            <div class="absolute left-full top-0 bg-white border border-gray-200 min-w-[210px] z-50 hidden group-hover/csr:block">
              <RouterLink to="/sustainability/csr/vision" class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100">Vision And Mission</RouterLink>
              <RouterLink to="/sustainability/csr/community" class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100">Local Community Development</RouterLink>
              <RouterLink to="/sustainability/csr/report" class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42]">CSR Report</RouterLink>
            </div>
          </div>
        </div>
      </div>

      <!-- Biodiversity -->
      <a href="#" class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap">
        BIODEVERSITY
      </a>

      <!-- Contact -->
      <a href="#" class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 hover:bg-white/20 transition whitespace-nowrap">
        CONTACT
      </a>

      <!-- Spacer -->
      <div class="flex-1"></div>

      <!-- Language -->
      <div class="relative group">
        <button class="text-white text-xs font-bold tracking-widest uppercase px-5 py-4 flex items-center gap-1.5 hover:bg-white/20 transition whitespace-nowrap">
          LANGUAGE <span class="text-[9px]">▾</span>
        </button>
        <div class="absolute top-full right-0 bg-white border border-gray-200 min-w-[140px] z-50 hidden group-hover:block">
          <a class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100 cursor-pointer">Indonesia</a>
          <a class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] border-b border-gray-100 cursor-pointer">English</a>
          <a class="block px-5 py-3 text-xs text-gray-700 hover:bg-green-50 hover:text-[#5F9E42] cursor-pointer">日本語</a>
        </div>
      </div>

    </div>
  </nav>
</template>

<script setup>
import { RouterLink } from 'vue-router'
</script>