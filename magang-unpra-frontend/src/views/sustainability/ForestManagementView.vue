<template>

  <!-- ── HERO ── -->
  <section class="relative h-[420px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('https://images.unsplash.com/photo-1448375240586-882707db888b?w=1920')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>

    <!-- Label atas kiri -->
    <div class="absolute top-10 left-10 flex items-center gap-2">
      <span class="w-1 h-5 bg-green-400 rounded-full inline-block"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>

    <!-- Teks bawah -->
    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4">
        Forest <span class="text-green-400">Management</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed">
        Sustainable and responsible forest management for a greener future in South Sumatra.
      </p>
      <!-- Breadcrumb -->
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400">
        <span class="hover:text-white cursor-pointer transition">Home</span>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">Forest Management</span>
      </div>
    </div>
  </section>

  <!-- ── STAT BAR ── -->
  <section class="bg-green-700 py-5">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-3 divide-x divide-green-600">
        <div class="text-center px-6">
          <p class="text-2xl font-black text-white">296,400 <span class="text-green-300 text-base font-semibold">ha</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">Concession Area</p>
        </div>
        <div class="text-center px-6">
          <p class="text-2xl font-black text-white">365 <span class="text-green-300 text-base font-semibold">days</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">Timber Delivery</p>
        </div>
        <div class="text-center px-6">
          <p class="text-2xl font-black text-white">Since <span class="text-green-300">1999</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">In Operation</p>
        </div>
      </div>
    </div>
  </section>

  <!-- ── INTRO ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-2 gap-16 items-center">

        <!-- Teks -->
        <div>
          <span class="inline-block bg-green-50 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-4">
            Our Approach
          </span>
          <h2 class="text-3xl font-black text-gray-900 leading-tight mb-5">
            High Quality Pulp from <span class="text-green-600">100% Sustainable Wood</span>
          </h2>
          <div class="space-y-4 text-gray-500 text-sm leading-relaxed">
            <p>
              PT Tanjungenim Lestari Pulp and Paper (TEL) produces high quality pulp from 100% wood
              based on industrial plantations sourced from a group company, PT Musi Hutan Persada (MHP),
              a highly dedicated Industrial Plantation Company with 296,400 ha of concession area in South Sumatra.
            </p>
            <p>
              MHP runs its operations with Sustainable and Sustainable Forest Management.
            </p>
          </div>

          <!-- Highlight box -->
          <div class="mt-6 border-l-4 border-green-500 pl-4 bg-green-50 py-3 pr-4 rounded-r-xl">
            <p class="text-green-800 text-sm font-medium leading-relaxed">
              TEL is working with MHP to develop new alternative raw materials — Eucalyptus Pellita Pulp
              has now become the main product.
            </p>
          </div>
        </div>

        <!-- Gambar -->
        <div class="relative">
          <div class="rounded-2xl overflow-hidden shadow-xl aspect-[4/3]">
            <img
              src="/images/hutan.jpg"
              alt="Hutan"
              class="w-full h-full object-cover"
            />
          </div>
          <!-- Badge tempel -->
          <div class="absolute -bottom-5 -left-5 bg-green-600 text-white rounded-xl px-5 py-3 shadow-lg">
            <p class="text-lg font-black">100%</p>
            <p class="text-green-200 text-xs">Sustainable Wood</p>
          </div>
        </div>

      </div>
    </div>
  </section>

  <!-- ── WOOD TYPES ── -->
  <section class="py-20 bg-gray-50">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-12">
        <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3">
          Raw Materials
        </span>
        <h2 class="text-3xl font-black text-gray-900">Types of Wood Used</h2>
        <p class="text-gray-400 text-sm mt-2 max-w-md mx-auto">
          Pulp is produced from two types of wood sourced from sustainable industrial plantations.
        </p>
      </div>

      <div class="grid grid-cols-2 gap-8">

        <!-- Acacia -->
        <div class="bg-white rounded-2xl overflow-hidden shadow-sm border border-gray-100 hover:shadow-md transition group">
          <div class="h-48 overflow-hidden">
            <img
  src="/images/accsia.jpg"
  alt="Acacia"
  class="w-full h-full object-cover"
/>
          </div>
          <div class="p-6">
            <div class="flex items-center gap-2 mb-2">
              <span class="w-2 h-2 rounded-full bg-green-500 inline-block"></span>
              <span class="text-xs font-bold text-green-600 tracking-widest uppercase">Initial Product</span>
            </div>
            <h3 class="text-lg font-black text-gray-900 mb-2">Acacia Mangium (AM)</h3>
            <p class="text-gray-400 text-sm leading-relaxed">
              The main product since the start of TEL's plant operation in 1999.
              Acacia Mangium Pulp served as the foundation of production.
            </p>
          </div>
        </div>

        <!-- Eucalyptus -->
        <div class="bg-white rounded-2xl overflow-hidden shadow-sm border border-gray-100 hover:shadow-md transition group">
          <div class="h-48 overflow-hidden">
             <img
              src="/images/eccalyptus.jpg"
              alt="Ecalyptus"
              class="w-full h-full object-cover"
            />
          </div>
          <div class="p-6">
            <div class="flex items-center gap-2 mb-2">
              <span class="w-2 h-2 rounded-full bg-emerald-500 inline-block"></span>
              <span class="text-xs font-bold text-emerald-600 tracking-widest uppercase">Main Product Now</span>
            </div>
            <h3 class="text-lg font-black text-gray-900 mb-2">Eucalyptus Pellita (EP)</h3>
            <p class="text-gray-400 text-sm leading-relaxed">
              Through continuous innovation and development, Eucalyptus Pellita Pulp
              has now become the primary raw material used by TEL.
            </p>
          </div>
        </div>

      </div>
    </div>
  </section>

  <!-- ── TIMBER DELIVERY ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-2 gap-16 items-center">

        <!-- Gambar kiri -->
        <div class="relative">
          <div class="rounded-2xl overflow-hidden shadow-xl aspect-[4/3]">
            <img
              src="https://images.unsplash.com/photo-1448375240586-882707db888b?w=800"
              alt="Timber Delivery"
              class="w-full h-full object-cover hover:scale-105 transition-transform duration-700"
            />
          </div>
          <!-- Floating card -->
          <div class="absolute -top-5 -right-5 bg-white border border-gray-100 shadow-lg rounded-xl px-5 py-4 w-44">
            <p class="text-3xl font-black text-green-600">365</p>
            <p class="text-gray-400 text-xs leading-tight mt-1">days/year timber delivery</p>
          </div>
        </div>

        <!-- Teks kanan -->
        <div>
          <span class="inline-block bg-green-50 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-4">
            Logistics
          </span>
          <h2 class="text-3xl font-black text-gray-900 leading-tight mb-5">
            Non-Stop Timber <span class="text-green-600">Delivery System</span>
          </h2>
          <p class="text-gray-500 text-sm leading-relaxed mb-6">
            Timber delivery from the Industrial Plantation area to the Factory Site continues
            365 days per year by truck through a special road network built by MHP.
            TEL sets standards for the size and quality of logs that must be met by MHP.
          </p>

          <!-- Checklist -->
          <ul class="space-y-3">
            <li v-for="item in deliveryFeatures" :key="item"
              class="flex items-start gap-3 text-sm text-gray-600">
              <span class="mt-0.5 w-5 h-5 rounded-full bg-green-100 text-green-600 flex items-center justify-center flex-shrink-0 text-xs font-bold">✓</span>
              {{ item }}
            </li>
          </ul>
        </div>

      </div>
    </div>
  </section>

  <!-- ── CTA ── -->
  <section class="py-16 bg-green-700">
    <div class="max-w-5xl mx-auto px-10 text-center">
      <h2 class="text-3xl font-black text-white mb-3">
        Committed to Sustainable Forest Management
      </h2>
      <p class="text-green-200 text-sm mb-8 max-w-lg mx-auto">
        TEL and MHP continuously strive to maintain the highest standards of environmental responsibility.
      </p>
      
    </div>
  </section>

</template>

<script setup>
const deliveryFeatures = [
  'Special road network built and maintained by MHP',
  'Strict log size and quality standards set by TEL',
  'Year-round continuous supply to factory site',
  'Efficient truck-based transportation system',
]
</script>