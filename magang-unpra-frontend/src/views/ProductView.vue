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
      <div class="relative overflow-hidden rounded-3xl bg-gray-100 mb-4 h-[600px] max-w-2xl mx-auto">
        <img
          v-for="(img, i) in images" :key="i"
          :src="img.url" :alt="img.caption"
          class="absolute inset-0 w-full h-full object-contain transition-opacity duration-500"
          :class="currentImage === i ? 'opacity-100' : 'opacity-0'"
        />
        <div class="absolute bottom-0 left-0 right-0 bg-black/50 text-white text-sm px-6 py-3">
          {{ images[currentImage].caption }}
        </div>
        <button @click="prevImage"
          class="absolute left-4 top-1/2 -translate-y-1/2 bg-black/50 text-white w-10 h-10 rounded-full text-xl flex items-center justify-center hover:bg-black/70 transition">‹</button>
        <button @click="nextImage"
          class="absolute right-4 top-1/2 -translate-y-1/2 bg-black/50 text-white w-10 h-10 rounded-full text-xl flex items-center justify-center hover:bg-black/70 transition">›</button>
      </div>

      <!-- Dots -->
      <div class="flex justify-center gap-2 mb-16">
        <button v-for="(_, i) in images" :key="i" @click="currentImage = i"
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

  <!-- Product Catalog dari Database -->
  <section class="py-20 bg-gradient-to-b from-white to-gray-50 border-t border-gray-100">
    <div class="max-w-6xl mx-auto px-6">
      <div class="text-center mb-14">
        <p class="text-green-600 text-xs font-semibold tracking-widest mb-3">OUR PRODUCTS</p>
        <h2 class="text-4xl font-black text-gray-900">Product Catalog</h2>
        <p class="text-gray-500 mt-3 max-w-xl mx-auto text-sm">
          High-quality pulp products manufactured with world-class technology and sustainable practices.
        </p>
      </div>

      <div v-if="loadingProducts" class="text-center py-10">
        <div class="w-8 h-8 border-4 border-green-600 border-t-transparent rounded-full animate-spin mx-auto"></div>
        <p class="text-gray-400 mt-3 text-sm">Loading products...</p>
      </div>

      <div v-else-if="products.length === 0" class="text-center py-16">
        <p class="text-5xl mb-4">📦</p>
        <p class="text-gray-400">Belum ada produk ditambahkan.</p>
      </div>

      <div v-else class="grid grid-cols-3 gap-8">
        <div v-for="product in products" :key="product.ID"
          class="group bg-white rounded-3xl overflow-hidden border border-gray-100 shadow-sm hover:shadow-xl transition-all duration-300 hover:-translate-y-1">
          <div class="relative h-56 bg-gradient-to-br from-green-50 to-gray-100 overflow-hidden">
            <img
              v-if="product.Images && product.Images.length > 0"
              :src="product.Images[0].image_url"
              :alt="product.name"
              class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
            />
            <div v-else class="w-full h-full flex flex-col items-center justify-center">
              <span class="text-6xl mb-2">🌿</span>
              <p class="text-green-400 text-xs font-semibold">TeL Product</p>
            </div>
            <div v-if="product.Images && product.Images.length > 1"
              class="absolute top-3 right-3 bg-black/60 text-white text-xs px-2 py-1 rounded-full">
              +{{ product.Images.length }} foto
            </div>
          </div>
          <div class="p-6">
            <h3 class="font-black text-gray-900 text-lg leading-tight mb-2">{{ product.name }}</h3>
            <p class="text-gray-500 text-sm leading-relaxed line-clamp-3 mb-5">{{ product.description }}</p>
            <div class="flex items-center justify-between pt-4 border-t border-gray-50">
              <span class="text-xs text-green-600 font-semibold bg-green-50 px-3 py-1 rounded-full">
                ✓ FSC Certified
              </span>
            </div>
          </div>
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
import { ref, onMounted } from 'vue'
import api from '../services/api'

// ===== CAROUSEL =====


const currentImage = ref(0)
const products = ref([])
const loadingProducts = ref(true)

const images = [
  {
    url: '/images/produk.jpeg',
    caption: 'TeL Pellita Bleached Kraft Pulp'
  },
  {
    url: '/images/produk2.jpeg',
    caption: 'Industrial pulp production facility'
  },
  {
    url: '/images/produk3.jpeg',
    caption: 'Sustainable plantation'
  },
  {
    url: '/images/produk.jpeg',
    caption: 'Forest management'
  },
  {
    url: '/images/produk2.jpeg',
    caption: 'Eucalyptus plantation'
  },
  {
    url: '/images/produk3.jpeg',
    caption: 'Acacia plantation'
  }
]

const prevImage = () => { currentImage.value = currentImage.value === 0 ? images.length - 1 : currentImage.value - 1 }
const nextImage = () => { currentImage.value = currentImage.value === images.length - 1 ? 0 : currentImage.value + 1 }

onMounted(async () => {
  try {
    const res = await api.get('/products')
    products.value = res.data.data
  } catch (err) {
    console.error(err)
  } finally {
    loadingProducts.value = false
  }
})

const specifications = [
  { parameter: 'Design Capacity Pulp Production', unit: 'Adt/year', value: '450,000', standard: 'BKP (Bleached Kraft Pulp)' },
  { parameter: 'Daily Production Capacity', unit: 'Adt/day', value: '1,430', standard: 'BKP' },
  { parameter: 'Chemical Recovery Capacity', unit: 'TDS/day', value: '2,400', standard: 'Solid burning in Recovery Boiler' },
  { parameter: 'Wood Species', unit: '-', value: 'Pellita (100%)', standard: 'Planted Plantation' },
  { parameter: 'Pulp Type', unit: '-', value: 'Bleached Kraft Pulp', standard: 'ECF Process' },
  { parameter: 'Brightness', unit: '% ISO', value: '89 - 90', standard: 'ISO 2470' },
  { parameter: 'Moisture Content', unit: '%', value: '≤ 10', standard: 'ISO 287' },
  { parameter: 'Mill Location', unit: '-', value: 'Banuayu, Muara Enim', standard: 'South Sumatra, Indonesia' },
  { parameter: 'Mill Area', unit: 'Hectare', value: '1,250', standard: 'Including plantation area' },
  { parameter: 'Employees', unit: 'People', value: '~1,600', standard: '~80% South Sumatra residents' },
]

const processSteps = [
  {
    icon: '🌳',
    image: '/images/accsia.jpg',
    title: 'Wood & Chip Handling',
    description: '...'
  },
  {
    icon: '♻️',
    image: '/images/accsia.jpg',
    title: 'Fiber Line',
    description: '...'
  }
]
</script>