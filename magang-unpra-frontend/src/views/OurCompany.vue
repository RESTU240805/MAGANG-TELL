<template>
  <!-- Hero / Company Profile Section -->
  <section class="relative min-h-screen flex items-center bg-cover bg-center"
    style="background-image: url('/images/gedung.jpeg')">
    <div class="absolute inset-0 bg-black/55"></div>
    <div class="relative z-10 w-full px-16 pt-32 pb-20 text-white">
      <h1 class="text-5xl font-black mb-6 italic" style="font-family: Georgia, serif;">Company Profile</h1>
      <div class="max-w-lg">
        <p class="leading-relaxed mb-4 text-sm italic font-medium">
          PT Tanjungenim Lestari Pulp and Paper (TeL) is a world-class manufacturer of high product
          quality and environmentally friendly market pulp mill.
          Established on June 18, 1990, TeL commenced construction in mid-1997 and commercial
          operations began in May 2000.
        </p>
        <p class="leading-relaxed mb-4 text-sm italic font-medium">
          The mill is strategically located in a 1,250 ha area in the Banuayu village, District Empat Petulai Dangku,
          Muara Enim Regency, South Sumatra province, Indonesia. As a Foreign Investment Company (PMA) of
          Marubeni Corporation, Japan, TeL is designated as a National Vital Objects Industrial sector (OVNI).
        </p>
        <p class="leading-relaxed text-sm italic font-medium">
          Presently, the mill maintains a market pulp production capacity of 490,000 Adt/year, supported by a
          dedicated workforce of 1,600 employees, with approximately 80% being residents of South Sumatra.
        </p>
      </div>
    </div>
  </section>

  <!-- Company Creed Section -->
<section class="py-16 bg-white">
  <div class="max-w-3xl mx-auto px-10 text-center">
    <h2 class="text-2xl font-black text-gray-900 mb-3">Company Creed</h2>
    <p class="text-gray-600 mb-2 text-sm italic">
      To Achieve sustainable growth in the pulp and paper industry...
    </p>
    <p class="text-green-600 font-bold text-sm mb-6">Creeds ( Values ) :</p>

    <!-- Gambar kaligrafi di tengah atas -->
    <div class="flex justify-center mb-6">
      <img src="/images/jepang.jpeg" alt="和 新 匠" class="h-36 w-auto object-contain" />
    </div>

    <!-- Kartu nilai horizontal -->
    <div class="grid grid-cols-3 gap-4">
      <div v-for="creed in creeds" :key="creed.title"
        class="border border-gray-200 rounded-xl p-4 text-left">
        <div class="flex items-start gap-2 mb-2">
          <div class="w-5 h-5 rounded-full border-2 border-green-600 flex-shrink-0 mt-0.5 flex items-center justify-center">
            <div class="w-2 h-2 bg-green-600 rounded-full"></div>
          </div>
          <div>
            <h3 class="font-black text-gray-900 text-sm">{{ creed.title }}</h3>
            <p class="text-green-600 text-xs italic">{{ creed.subtitle }}</p>
          </div>
        </div>
        <p class="text-gray-600 text-xs leading-relaxed">- {{ creed.description }}</p>
      </div>
    </div>
  </div>
</section>

 <!-- PEFC Section -->
<section class="py-12 bg-gray-50">
  <div class="max-w-3xl mx-auto px-10">
    <div class="bg-white rounded-xl border border-gray-200 overflow-hidden flex">

      <div class="w-36 flex-shrink-0 bg-gray-100 border-r border-gray-200 flex items-center justify-center p-4">
        <img src="/images/pdf.png" alt="PEFC Document" class="w-24 h-auto object-contain" />
      </div>

      <div class="flex flex-col justify-center gap-2.5 px-6 py-5">
        <span class="text-[11px] font-medium text-green-700 bg-green-50 px-2.5 py-0.5 rounded-full w-fit">DOCUMENT PREVIEW</span>

        <div>
          <p class="text-sm font-medium text-gray-900 mb-1">PEFC Chain of Custody Commitment Statement</p>
          <div class="flex items-center gap-4 flex-wrap">
            <span class="text-xs text-gray-500 flex items-center gap-1">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5"><rect x="3" y="4" width="18" height="18" rx="2"/><line x1="16" y1="2" x2="16" y2="6"/><line x1="8" y1="2" x2="8" y2="6"/><line x1="3" y1="10" x2="21" y2="10"/></svg>
              April 2023
            </span>
            <span class="text-xs text-gray-500 flex items-center gap-1">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-3-3v6M5 8l1.5-1.5A2 2 0 018 6h8a2 2 0 012 2v10a2 2 0 01-2 2H8a2 2 0 01-2-2V8z"/></svg>
              PDF · 1.2 MB
            </span>
          </div>
        </div>

        <p class="text-xs text-gray-500 leading-relaxed">
          PT. Tanjungenim Lestari Pulp and Paper is committed to sourcing wood-based products in an environmentally and socially responsible manner, in full compliance with EU Timber Regulation and sustainable forest management practices.
        </p>

        <div class="flex gap-2 flex-wrap mt-1">
          <a href="/files/PEFC_Statement.pdf" download="PEFC-Chain-of-Custody.pdf" class="inline-flex items-center gap-1.5 bg-green-700 hover:bg-green-800 text-white text-xs font-medium px-4 py-1.5 rounded-md transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/></svg>
            Download PDF
          </a>
          <button @click="showPdf = true" class="inline-flex items-center gap-1.5 border border-gray-300 hover:bg-gray-50 text-gray-600 text-xs font-medium px-4 py-1.5 rounded-md transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/></svg>
            Quick View
          </button>
        </div>
      </div>

    </div>
  </div>
</section>

<!-- Modal PDF Viewer -->
<Teleport to="body">
  <div v-if="showPdf" class="fixed inset-0 z-50 bg-black/70 flex flex-col">
    <div class="flex items-center gap-3 px-4 py-3 bg-gray-900">
      <button @click="showPdf = false" class="flex items-center gap-2 text-white text-sm hover:text-green-400 transition-colors">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M10 19l-7-7m0 0l7-7m-7 7h18"/>
        </svg>
        Back
      </button>
      <span class="text-gray-400 text-sm">|</span>
      <span class="text-white text-sm font-medium">PEFC Chain of Custody Commitment Statement</span>
      <a href="/files/PEFC_Statement.pdf" download="PEFC-Chain-of-Custody.pdf" class="ml-auto flex items-center gap-1.5 bg-green-700 hover:bg-green-800 text-white text-xs font-medium px-3 py-1.5 rounded-md transition-colors">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"/></svg>
        Download
      </a>
    </div>
    <iframe src="/files/PEFC_Statement.pdf" class="flex-1 w-full" frameborder="0"></iframe>
  </div>
</Teleport>

  <!-- Footer -->
  <section class="py-4 text-white text-center" style="background-color: #5F9E42;">
    <p class="text-sm italic">Copyright 2019 PT TELPP. All right reserved.</p>
  </section>
</template>

<script setup>

import { ref } from 'vue'
const showPdf = ref(false)

const creeds = [
  {
    title: 'Fairness',
    subtitle: 'To be fair and decent.',
    description: 'We shall comply with the laws and follow fair corporate practices.'
  },
  {
    title: 'Innovation',
    subtitle: 'To be active and innovative.',
    description: 'We shall constantly strive hard to improve our performance.'
  },
  {
    title: 'Harmony',
    subtitle: 'To respect each other and cooperate.',
    description: 'We shall stay in touch with society and stakeholders by engaging in corporate activities which advance our credibility as preferred principals.'
  },
]
</script>