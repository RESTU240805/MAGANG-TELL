<template>
  <!-- Hero Section -->
  <section class="relative h-screen flex items-center justify-start bg-cover bg-center"
    style="background-image: url('https://via.placeholder.com/1920x1080/2d4a2d/ffffff?text=Factory+Image')">
    <div class="absolute inset-0 bg-black/50"></div>
    <div class="relative z-10 max-w-7xl mx-auto px-10">
      <span class="text-xs border border-white text-white px-3 py-1 rounded-full mb-6 inline-block">
        FSC & PEFC Certified Production
      </span>
      <h1 class="text-6xl font-black text-white leading-tight mt-4">
        Sustaining<br>Nature,<br>
        <span class="text-green-400">Crafting Paper.</span>
      </h1>
      <p class="text-gray-200 mt-6 max-w-lg text-lg">
        Leading the global transition towards sustainable pulp production
        and responsible forest management in North Sumatra.
      </p>
    </div>
  </section>

  <!-- Legacy Section -->
  <section class="py-24 bg-white">
    <div class="max-w-7xl mx-auto px-10 flex gap-16 items-start">
      <div class="min-w-fit">
        <p class="text-green-600 text-xs font-semibold tracking-widest flex items-center gap-2">
          <span class="w-2 h-2 bg-green-600 rounded-full"></span> SECTION 01
        </p>
        <p class="text-xs tracking-widest text-gray-400 mt-1 border-l-2 border-gray-300 pl-2">
          LEGACY &<br>HISTORY
        </p>
      </div>
      <div>
        <h2 class="text-5xl font-black text-gray-900 leading-tight mb-6">
          World-Class Pulp<br>Manufacturer
        </h2>
        <p class="text-gray-600 leading-relaxed mb-4">
          PT Tanjungenim Lestari Pulp and Paper (TeL) is world class manufacturer of high product quality
          and environmental friendly market pulp mill. This was established on June 18, 1990, commenced
          construction in mid-1997 and the commercial operation started on May, 2000.
        </p>
        <p class="text-gray-600 leading-relaxed">
          TeL is a Foreign Investment Company (PMA)- Marubeni Corporation, Japan, as the National Vital
          Objects Industrial sector (OVNI) declared by the Minister of Industry in 2014. The mill has market
          pulp production capacity of 490,000 Adt/year. Presently mill has 1600 employees and support
          workforce together where ~80% of them are residents of South Sumatra.
        </p>
      </div>
    </div>
  </section>

  <!-- Creed Section -->
  <section class="py-24 bg-gray-50">
    <div class="max-w-7xl mx-auto px-10">
      <div class="flex gap-16 items-start mb-12">
        <div class="min-w-fit">
          <p class="text-green-600 text-xs font-semibold tracking-widest flex items-center gap-2">
            <span class="w-2 h-2 bg-green-600 rounded-full"></span> SECTION 02
          </p>
          <p class="text-xs tracking-widest text-gray-400 mt-1 border-l-2 border-gray-300 pl-2">
            THE CREED
          </p>
        </div>
        <div>
          <div class="bg-black text-white px-6 py-3 inline-block rounded mb-4 text-2xl font-bold">
            和 新 匠
          </div>
          <p class="text-gray-600 italic max-w-xl">
            "To Achieve sustainable growth in the pulp and paper industry and to operate in harmony
            with all stake holder and environment for creating long term prosperity and better quality of life."
          </p>
        </div>
      </div>

      <!-- Values -->
      <div class="grid grid-cols-3 gap-6 mt-8">
        <div class="bg-white rounded-2xl p-8 border border-gray-100 shadow-sm">
          <div class="text-green-600 text-3xl mb-4">⚖️</div>
          <h3 class="text-xl font-bold text-gray-900 mb-1">Fairness</h3>
          <p class="text-green-600 text-xs font-semibold tracking-widest mb-3">TO BE FAIR AND DECENT.</p>
          <p class="text-gray-500 text-sm">We shall comply with the laws and follow fair corporate practices.</p>
        </div>
        <div class="bg-white rounded-2xl p-8 border border-gray-100 shadow-sm">
          <div class="text-green-600 text-3xl mb-4">✦</div>
          <h3 class="text-xl font-bold text-gray-900 mb-1">Innovation</h3>
          <p class="text-green-600 text-xs font-semibold tracking-widest mb-3">TO BE ACTIVE AND INNOVATIVE.</p>
          <p class="text-gray-500 text-sm">We shall constantly strive hard to improve our performance.</p>
        </div>
        <div class="bg-white rounded-2xl p-8 border border-gray-100 shadow-sm">
          <div class="text-green-600 text-3xl mb-4">☯️</div>
          <h3 class="text-xl font-bold text-gray-900 mb-1">Harmony</h3>
          <p class="text-green-600 text-xs font-semibold tracking-widest mb-3">TO RESPECT EACH OTHER AND COOPERATE.</p>
          <p class="text-gray-500 text-sm">We shall stay in touch with society and stakeholders by engaging in corporate activities.</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Latest Insights -->
  <section class="py-24 bg-white">
    <div class="max-w-7xl mx-auto px-10">
      <div class="flex justify-between items-start mb-12">
        <div>
          <p class="text-orange-500 text-xs font-semibold tracking-widest mb-2">STRATEGIC MEDIA</p>
          <h2 class="text-4xl font-black text-gray-900">Latest Insights</h2>
          <p class="text-gray-500 mt-2 max-w-md">Inside look at our latest reforestation projects, industrial innovations, and community impact.</p>
        </div>
        <button class="border border-gray-300 px-5 py-2 rounded-full text-sm hover:bg-gray-50 transition">
          View More →
        </button>
      </div>

      <div class="grid grid-cols-3 gap-6">
        <div class="rounded-2xl overflow-hidden border border-gray-100 shadow-sm">
          <div class="h-48 bg-gray-200"></div>
          <div class="p-6">
            <p class="text-green-600 text-xs font-semibold tracking-widest mb-2">OPERATIONAL EXCELLENCE</p>
            <h3 class="font-bold text-gray-900 mb-2">TeLpp Achieves PEFC Re-certification</h3>
            <p class="text-gray-500 text-sm mb-4">Our commitment to sustainable forest management is reaffirmed with the successful PEFC audit.</p>
            <button class="bg-orange-500 text-white px-4 py-2 rounded-full text-sm hover:bg-orange-600 transition">Read More</button>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- Footer -->
  <footer class="bg-gray-900 text-white py-16">
    <div class="max-w-7xl mx-auto px-10 grid grid-cols-4 gap-10">
      <div>
        <div class="flex items-center gap-2 mb-4">
          <div class="w-8 h-8 bg-green-600 rounded-lg"></div>
          <span class="text-xl font-black">TeLpp.</span>
        </div>
        <p class="text-gray-400 text-sm leading-relaxed">
          Global leaders in high-performance, sustainable pulp solutions. Rooted in North Sumatra, driven by innovation for a resilient future.
        </p>
      </div>
      <div>
        <p class="text-xs font-semibold tracking-widest text-gray-500 mb-4">HEADQUARTERS</p>
        <p class="text-sm text-gray-300">📍 Parmeslan, Tobe District<br>North Sumatra, Indonesia</p>
        <p class="text-sm text-gray-300 mt-2">📞 +62 (823) 123-4567</p>
        <p class="text-sm text-gray-300 mt-2">✉️ corporate@telpp.com</p>
      </div>
      <div>
        <p class="text-xs font-semibold tracking-widest text-gray-500 mb-4">INDUSTRIAL STRATEGY</p>
        <ul class="space-y-2 text-sm text-gray-300">
          <li class="hover:text-green-400 cursor-pointer">Circular Economy</li>
          <li class="hover:text-green-400 cursor-pointer">Sustainability Roadmap</li>
          <li class="hover:text-green-400 cursor-pointer">Investor Relations</li>
          <li class="hover:text-green-400 cursor-pointer">Media Assets</li>
          <li class="hover:text-green-400 cursor-pointer">Privacy Ethics</li>
        </ul>
      </div>
      <div class="bg-green-900 rounded-2xl p-6">
        <p class="text-orange-400 text-xs font-semibold tracking-widest mb-2">OUR CONCESSION</p>
        <p class="font-bold text-lg">Preserving Biodiversity in North Sumatra.</p>
      </div>
    </div>
    <div class="max-w-7xl mx-auto px-10 mt-10 pt-6 border-t border-gray-800 flex justify-between text-xs text-gray-500">
      <span>© 2024 Tero Pulp Lestari. All Rights Reserved.</span>
      <div class="flex gap-4">
        <span>FSC Verified</span>
        <span>PEFC Certified</span>
        <span>ISO 14001</span>
      </div>
    </div>
  </footer>
</template>