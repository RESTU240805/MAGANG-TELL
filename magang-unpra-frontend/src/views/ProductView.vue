<template>
  <!-- Hero -->
  <section class="relative h-64 flex items-center justify-center bg-cover bg-center"
    style="background-image: url('https://images.unsplash.com/photo-1504328345606-18bbc8c9d7d1?w=1920')">
    <div class="absolute inset-0 bg-black/60"></div>
    <div class="relative z-10 text-center text-white">
      <p class="text-sm text-gray-300 mb-2">You are here: <span class="text-white font-medium">Home » Products</span></p>
      <h1 class="text-4xl font-black tracking-wide">PRODUCTS</h1>
    </div>
  </section>

  <!-- Product Intro -->
  <section class="py-16 bg-white">
    <div class="max-w-5xl mx-auto px-6">
      <h2 class="text-3xl font-black text-green-600 mb-6">Products</h2>
      <p class="text-gray-700 mb-3">
        Our main product is <strong class="text-green-600">TeL Pellita Bleached Kraft Pulp</strong> based on 100% planted Pellita trees.
      </p>
      <p class="text-gray-600 mb-3">
        TeL focuses on product quality in totality through tight quality control systems and also covered under QMS ISO 9001.
        The continual improvement for quality and process is important drive at TeL and for this we follow proactive approach
        to interact with our customers.
      </p>
      <p class="text-gray-600 mb-10">
        Based on end products, TeL pulp mostly used as raw material of Tissue, Coating base papers, Wood Free paper &
        paperboard, and others specialty grades.
      </p>

      <!-- Image Carousel -->
      <div class="relative overflow-hidden rounded-2xl bg-gray-100 mb-4 h-96">
        <img
          v-for="(img, i) in images"
          :key="i"
          :src="img.url"
          :alt="img.caption"
          class="absolute inset-0 w-full h-full object-cover transition-opacity duration-500"
          :class="currentImage === i ? 'opacity-100' : 'opacity-0'"
        />

        <!-- Caption -->
        <div class="absolute bottom-0 left-0 right-0 bg-black/50 text-white text-sm px-6 py-3">
          {{ images[currentImage].caption }}
        </div>

        <!-- Arrows -->
        <button @click="prevImage"
          class="absolute left-4 top-1/2 -translate-y-1/2 bg-black/50 text-white w-10 h-10 rounded-full text-xl flex items-center justify-center hover:bg-black/70 transition">
          ‹
        </button>
        <button @click="nextImage"
          class="absolute right-4 top-1/2 -translate-y-1/2 bg-black/50 text-white w-10 h-10 rounded-full text-xl flex items-center justify-center hover:bg-black/70 transition">
          ›
        </button>
      </div>

      <!-- Dots -->
      <div class="flex justify-center gap-2 mb-16">
        <button
          v-for="(_, i) in images" :key="i"
          @click="currentImage = i"
          :class="['w-2.5 h-2.5 rounded-full transition', currentImage === i ? 'bg-green-600 scale-125' : 'bg-gray-300']">
        </button>
      </div>

      <!-- Product Specs -->
      <div class="grid grid-cols-3 gap-6 mb-4">
        <div class="bg-green-50 rounded-2xl p-6 border border-green-100 text-center">
          <p class="text-3xl font-black text-green-600">490,000</p>
          <p class="text-gray-600 text-sm mt-1">Adt/year Production Capacity</p>
        </div>
        <div class="bg-green-50 rounded-2xl p-6 border border-green-100 text-center">
          <p class="text-3xl font-black text-green-600">ISO 9001</p>
          <p class="text-gray-600 text-sm mt-1">Quality Management Certified</p>
        </div>
        <div class="bg-green-50 rounded-2xl p-6 border border-green-100 text-center">
          <p class="text-3xl font-black text-green-600">100%</p>
          <p class="text-gray-600 text-sm mt-1">Planted Pellita Trees</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Pulp Production Process -->
  <section class="py-16 bg-gray-50">
    <div class="max-w-5xl mx-auto px-6">
      <div class="text-center mb-12">
        <p class="text-green-600 text-xs font-semibold tracking-widest mb-2">HOW WE MAKE IT</p>
        <h2 class="text-4xl font-black text-gray-900">Pulp Production Process</h2>
        <p class="text-gray-500 mt-3 max-w-2xl mx-auto">
          From forest to finished product, every step follows strict environmental and quality standards.
        </p>
      </div>

      <div class="space-y-8">
        <div v-for="(step, index) in processSteps" :key="index"
          class="bg-white rounded-2xl overflow-hidden shadow-sm border border-gray-100 flex flex-col md:flex-row"
          :class="index % 2 === 1 ? 'md:flex-row-reverse' : ''">

          <!-- Image -->
          <div class="md:w-2/5 h-56 md:h-auto overflow-hidden flex-shrink-0">
            <img :src="step.image" :alt="step.title" class="w-full h-full object-cover" />
          </div>

          <!-- Content -->
          <div class="flex-1 p-8 flex flex-col justify-center">
            <div class="flex items-center gap-3 mb-3">
              <div class="w-10 h-10 rounded-xl flex items-center justify-center text-xl flex-shrink-0"
                :class="step.bgColor">
                {{ step.icon }}
              </div>
              <span class="text-xs font-bold text-gray-400 tracking-widest">STEP {{ index + 1 }}</span>
            </div>
            <h3 class="text-xl font-black text-gray-900 mb-3">{{ step.title }}</h3>
            <p class="text-gray-500 leading-relaxed text-sm">{{ step.description }}</p>
          </div>
        </div>
      </div>
    </div>
  </section>

  <!-- End Uses -->
  <section class="py-16 bg-white">
    <div class="max-w-5xl mx-auto px-6">
      <div class="text-center mb-10">
        <p class="text-green-600 text-xs font-semibold tracking-widest mb-2">APPLICATIONS</p>
        <h2 class="text-3xl font-black text-gray-900">End Product Uses</h2>
        <p class="text-gray-500 mt-2">TeL pulp is used across a wide range of paper and packaging applications worldwide.</p>
      </div>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
        <div v-for="use in endUses" :key="use.name"
          class="text-center p-6 rounded-2xl bg-gray-50 border border-gray-100 hover:border-green-300 hover:bg-green-50 transition group">
          <div class="text-4xl mb-3">{{ use.icon }}</div>
          <p class="font-bold text-gray-800 text-sm group-hover:text-green-700 transition">{{ use.name }}</p>
          <p class="text-gray-400 text-xs mt-1">{{ use.desc }}</p>
        </div>
      </div>
    </div>
  </section>

  <!-- Certifications -->
  <section class="py-16 bg-gray-900 text-white">
    <div class="max-w-5xl mx-auto px-6 text-center">
      <p class="text-green-400 text-xs font-semibold tracking-widest mb-2">VERIFIED & CERTIFIED</p>
      <h2 class="text-3xl font-black mb-10">Our Certifications</h2>
      <div class="grid grid-cols-3 gap-6">
        <div class="bg-gray-800 rounded-2xl p-6 border border-gray-700">
          <p class="text-2xl font-black text-green-400 mb-2">FSC®</p>
          <p class="text-sm text-gray-400">Forest Stewardship Council certified sustainable forest management</p>
        </div>
        <div class="bg-gray-800 rounded-2xl p-6 border border-gray-700">
          <p class="text-2xl font-black text-green-400 mb-2">PEFC™</p>
          <p class="text-sm text-gray-400">Programme for the Endorsement of Forest Certification</p>
        </div>
        <div class="bg-gray-800 rounded-2xl p-6 border border-gray-700">
          <p class="text-2xl font-black text-green-400 mb-2">ISO 9001</p>
          <p class="text-sm text-gray-400">Quality Management System for consistent product quality</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref } from 'vue'

const currentImage = ref(0)

const images = [
  {
    url: 'https://images.unsplash.com/photo-1504328345606-18bbc8c9d7d1?w=1200',
    caption: 'TeL Pellita Bleached Kraft Pulp — Ready for shipment'
  },
  {
    url: 'https://images.unsplash.com/photo-1565193566173-7a0ee3dbe261?w=1200',
    caption: 'Industrial pulp production facility at PT. TeLPP'
  },
  {
    url: 'https://images.unsplash.com/photo-1473341304170-971dccb5ac1e?w=1200',
    caption: 'Sustainable plantation — 100% Pellita trees'
  },
  {
    url: 'https://images.unsplash.com/photo-1530973428-5bf2db2e4d71?w=1200',
    caption: 'Advanced bleaching process — ECF technology'
  },
  {
    url: 'https://images.unsplash.com/photo-1518770660439-4636190af475?w=1200',
    caption: 'Quality control and testing laboratory'
  },
  {
    url: 'https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?w=1200',
    caption: 'Pulp bales ready for global export'
  },
  {
    url: 'https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?w=1200',
    caption: 'Modern production equipment and automation systems'
  },
]

const prevImage = () => {
  currentImage.value = currentImage.value === 0 ? images.length - 1 : currentImage.value - 1
}
const nextImage = () => {
  currentImage.value = currentImage.value === images.length - 1 ? 0 : currentImage.value + 1
}

const processSteps = [
  {
    icon: '🌳',
    bgColor: 'bg-green-100',
    image: 'https://images.unsplash.com/photo-1448375240586-882707db888b?w=800',
    title: 'Wood Harvesting & Chipping',
    description: 'Pellita trees grown in our sustainable plantation are harvested after reaching maturity (5-7 years). The logs are debarked and chipped into uniform wood chips of approximately 25mm × 25mm for optimal cooking efficiency. Only sustainably sourced wood from certified plantations is used.'
  },
  {
    icon: '♻️',
    bgColor: 'bg-blue-100',
    image: 'https://images.unsplash.com/photo-1565193566173-7a0ee3dbe261?w=800',
    title: 'Kraft Cooking (Digester)',
    description: 'Wood chips are fed into large digesters and cooked with white liquor (sodium hydroxide and sodium sulfide) at high temperature and pressure. This Kraft (sulfate) process dissolves the lignin binding the cellulose fibers, separating them into pulp while preserving superior fiber strength.'
  },
  {
    icon: '🔍',
    bgColor: 'bg-yellow-100',
    image: 'https://images.unsplash.com/photo-1504328345606-18bbc8c9d7d1?w=800',
    title: 'Washing & Screening',
    description: 'The cooked brown stock pulp is thoroughly washed to remove spent cooking chemicals (black liquor) which are sent to the chemical recovery system for recycling. Screening equipment removes oversized fiber bundles and contaminants to ensure uniform fiber quality throughout the process.'
  },
  {
    icon: '⚗️',
    bgColor: 'bg-purple-100',
    image: 'https://images.unsplash.com/photo-1518770660439-4636190af475?w=800',
    title: 'Oxygen Delignification',
    description: 'Before bleaching, the pulp undergoes oxygen delignification using oxygen and caustic soda to remove residual lignin. This pre-bleaching stage significantly reduces the chemical load on the bleaching plant and minimizes environmental impact by lowering chlorine compound usage in subsequent stages.'
  },
  {
    icon: '✨',
    bgColor: 'bg-cyan-100',
    image: 'https://images.unsplash.com/photo-1530973428-5bf2db2e4d71?w=800',
    title: 'ECF Bleaching',
    description: 'TeL uses Elemental Chlorine Free (ECF) bleaching with a multi-stage sequence using chlorine dioxide (ClO₂), sodium hydroxide, and hydrogen peroxide. This achieves high brightness (ISO 89-90%) while maintaining fiber strength and meeting strict environmental discharge standards for effluent treatment.'
  },
  {
    icon: '🏭',
    bgColor: 'bg-orange-100',
    image: 'https://images.unsplash.com/photo-1473341304170-971dccb5ac1e?w=800',
    title: 'Sheet Forming & Drying',
    description: 'The pulp slurry is spread onto a forming wire where water drains through gravity and vacuum suction. The wet sheet passes through press sections for mechanical water removal, then through steam-heated dryer cylinders that reduce moisture content to approximately 10% for optimal product quality.'
  },
  {
    icon: '🧪',
    bgColor: 'bg-red-100',
    image: 'https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?w=800',
    title: 'Quality Control & Testing',
    description: 'Every batch undergoes rigorous quality testing in our ISO-certified laboratory. Parameters tested include brightness, viscosity, tensile strength, tear resistance, drainage, and fiber length distribution. Only pulp meeting our strict specifications is approved for baling and shipment to customers worldwide.'
  },
  {
    icon: '📦',
    bgColor: 'bg-green-100',
    image: 'https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?w=800',
    title: 'Cutting, Baling & Shipping',
    description: 'The dried pulp sheet is cut into standard dimensions and compressed into bales of approximately 250 kg each. Bales are wrapped, strapped, and labeled with product specifications, batch number, and quality test results. Products are then loaded into containers for global shipment to customers in Asia, Europe, and beyond.'
  },
]

const endUses = [
  { icon: '🧻', name: 'Tissue Paper', desc: 'Facial tissue, toilet tissue, paper towels' },
  { icon: '📄', name: 'Wood Free Paper', desc: 'Office paper, copy paper, printing paper' },
  { icon: '📦', name: 'Paperboard', desc: 'Packaging boards, carton boxes' },
  { icon: '🎨', name: 'Coating Base', desc: 'Coated papers, magazine papers' },
]
</script>