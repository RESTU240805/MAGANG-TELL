<template>
  <!-- Hero -->
  <section class="relative h-64 flex items-center justify-center bg-cover bg-center"
    style="background-image: url('https://images.unsplash.com/photo-1437622368342-7a3d73a34c8f?w=1920')">
    <div class="absolute inset-0 bg-black/60"></div>
    <div class="relative z-10 text-center text-white">
      <p class="text-sm text-gray-300 mb-2">You are here: <span class="text-white font-medium">Home » Biodiversity</span></p>
      <h1 class="text-4xl font-black tracking-wide">BIODIVERSITY</h1>
    </div>
  </section>

  <!-- Biodiversity Section -->
  <section class="py-24 bg-white overflow-hidden">
    <div class="max-w-7xl mx-auto px-6">
      <div class="bg-gray-900 rounded-3xl p-10 md:p-20 flex flex-col lg:flex-row items-center gap-16 relative shadow-2xl">
        <div class="lg:w-1/2 relative z-10">
          <div class="inline-block px-4 py-1 rounded-full bg-green-600/20 text-green-400 text-sm font-bold uppercase tracking-widest mb-6">
            Conservation First
          </div>
          <h2 class="text-4xl md:text-5xl font-black text-white mb-8 leading-tight">
            Our Dedication to<br/>
            <span class="text-green-400">Biodiversity</span>
          </h2>
          <p class="text-white/70 text-lg mb-10 leading-relaxed">
            We don't just harvest; we protect. Our forest concessions include 20% dedicated conservation zones
            that serve as safe havens for endangered species and natural ecosystems.
          </p>
          <ul class="grid grid-cols-2 gap-6 mb-10">
            <li v-for="item in features" :key="item" class="flex items-center gap-3 text-white font-bold">
              <span class="text-green-400 bg-white/10 p-2 rounded-xl text-lg">✓</span>
              {{ item }}
            </li>
          </ul>
        </div>
        <div class="lg:w-1/2 w-full h-96 lg:h-[500px] rounded-3xl overflow-hidden shadow-2xl relative group">
          <img
            src="https://images.unsplash.com/photo-1437622368342-7a3d73a34c8f?w=1920"
            alt="Biodiversity"
            class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-1000"
          />
        </div>
      </div>
    </div>
  </section>

  <!-- Stats -->
  <section class="py-16 bg-gray-50">
    <div class="max-w-5xl mx-auto px-6">
      <div class="grid grid-cols-4 gap-6 text-center">
        <div v-for="stat in stats" :key="stat.label"
          class="bg-white rounded-2xl p-6 border border-gray-100 shadow-sm">
          <p class="text-3xl font-black text-green-600 mb-2">{{ stat.value }}</p>
          <p class="text-gray-500 text-sm">{{ stat.label }}</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Conservation Programs -->
  <section class="py-16 bg-white">
    <div class="max-w-5xl mx-auto px-6">
      <div class="text-center mb-12">
        <p class="text-green-600 text-xs font-semibold tracking-widest mb-2">OUR PROGRAMS</p>
        <h2 class="text-3xl font-black text-gray-900">Conservation Programs</h2>
      </div>
      <div class="grid grid-cols-2 gap-6">
        <div v-for="program in programs" :key="program.title"
          class="bg-gray-50 rounded-2xl p-8 border border-gray-100 hover:border-green-200 hover:shadow-md transition">
          <div class="text-4xl mb-4">{{ program.icon }}</div>
          <h3 class="font-black text-gray-900 text-lg mb-2">{{ program.title }}</h3>
          <p class="text-gray-500 text-sm leading-relaxed">{{ program.description }}</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Kehati Embed Section -->
  <section class="py-16 bg-gray-50">
    <div class="max-w-7xl mx-auto px-6">
      <div class="text-center mb-8">
        <p class="text-green-600 text-xs font-semibold tracking-widest mb-2">KEHATI PROGRAM</p>
        <h2 class="text-3xl font-black text-gray-900">Keanekaragaman Hayati</h2>
        <p class="text-gray-500 mt-2 text-sm">Program pelestarian keanekaragaman hayati PT. TeLPP</p>
      </div>

      <!-- Iframe Kehati -->
      <div class="rounded-2xl overflow-hidden shadow-xl border border-gray-200" style="height: 800px;">
        <iframe
          src="https://kehati.telpp.com/"
          class="w-full h-full border-0"
          title="Kehati - PT Tanjungenim Lestari Pulp and Paper"
        ></iframe>
      </div>

      <!-- Fallback jika iframe diblokir -->
      <div class="mt-4 text-center">
        <p class="text-gray-400 text-xs">Jika konten tidak tampil,
          <a href="https://kehati.telpp.com/" target="_blank"
            class="text-green-600 font-semibold hover:underline">
            klik di sini untuk membuka halaman Kehati
          </a>
        </p>
      </div>
    </div>
  </section>
</template>

<script setup>
const features = [
  'Wildlife Corridors',
  'Native Restoration',
  'Biodiversity Audits',
  'Species Monitoring',
]

const stats = [
  { value: '20%', label: 'Conservation Zones' },
  { value: '50+', label: 'Protected Species' },
  { value: '1,250', label: 'Hectares Protected' },
  { value: '25+', label: 'Years of Conservation' },
]

const programs = [
  {
    icon: '🦜',
    title: 'Wildlife Corridor Program',
    description: 'Creating connected pathways through our plantation areas to allow wildlife to move freely between natural forest patches, reducing habitat fragmentation.'
  },
  {
    icon: '🌿',
    title: 'Native Species Restoration',
    description: 'Replanting native tree species in degraded areas within our concession to restore natural ecosystem functions and increase biodiversity.'
  },
  {
    icon: '🔬',
    title: 'Biodiversity Monitoring',
    description: 'Regular biodiversity audits and species monitoring programs conducted in collaboration with research institutions and conservation organizations.'
  },
  {
    icon: '🐾',
    title: 'Endangered Species Protection',
    description: 'Special protection measures for endangered species found within our concession area, including Sumatran tigers, elephants, and orangutans.'
  },
]
</script>