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
  <!-- PEFC Section -->
<section class="py-12 bg-gray-50">
  <div class="max-w-3xl mx-auto px-10">
    <div class="bg-white rounded-xl border border-gray-200 shadow-sm overflow-hidden flex">
      <div class="w-48 flex-shrink-0 border-r border-gray-200 overflow-hidden">
        <img
          src="/images/pdf.png"
          alt="PEFC Document"
          class="w-full h-full object-cover object-top"
        />
      </div>
      <div class="flex gap-3">
  <a href="/files/PEFC_Statement.pdf" download="PEFC-Chain-of-Custody.pdf" class="bg-green-600 text-white text-xs px-5 py-2 rounded hover:bg-green-700 transition">↓ Download PDF</a>
  <a href="/files/PEFC_Statement.pdf" target="_blank" rel="noopener noreferrer" class="border border-gray-300 text-gray-600 text-xs px-5 py-2 rounded hover:bg-gray-50 transition">👁 Quick View</a>
</div>
    </div>
  </div>
</section>

  <!-- Footer -->
  <section class="py-4 text-white text-center" style="background-color: #5F9E42;">
    <p class="text-sm italic">Copyright 2019 PT TELPP. All right reserved.</p>
  </section>
</template>

<script setup>
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