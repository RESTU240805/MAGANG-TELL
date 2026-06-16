<template>
  <div class="min-h-screen bg-slate-50 text-slate-900 font-sans antialiased selection:bg-emerald-100">

    <!-- ─── HERO ──────────────────────────────────────────────── -->
    <header class="relative h-[85vh] flex items-center justify-center bg-zinc-900 overflow-hidden">
      <div class="absolute inset-0 z-0">
        <img
          src="/images/gedung.jpeg"
          class="w-full h-full object-cover opacity-60 animate-subtle-zoom"
          alt="Hero Background"
          @error="(e) => e.target.src='https://images.unsplash.com/photo-1541888946425-d81bb19240f5?q=80&w=1600'"
        />
        <div class="absolute inset-0 bg-gradient-to-b from-black/60 via-black/20 to-slate-50"></div>
      </div>

      <div class="container mx-auto px-6 lg:px-16 relative z-10 text-center">
        <Transition name="reveal-down" appear>
          <div class="inline-flex items-center gap-2 px-5 py-2.5 rounded-full bg-emerald-500/20 backdrop-blur-md border border-emerald-500/30 text-emerald-400 text-sm font-bold uppercase tracking-[0.2em] mb-8">
            <span class="relative flex h-2 w-2">
              <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-emerald-400 opacity-75"></span>
              <span class="relative inline-flex rounded-full h-2 w-2 bg-emerald-500"></span>
            </span>
            PRODUCT CATALOG
          </div>
        </Transition>

        <Transition name="reveal-up" appear>
          <h1 class="text-6xl md:text-8xl font-serif font-bold text-white mb-6 leading-[1.1]">
            Katalog Produk<br>
            <span class="text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-teal-200">
              Unggulan
            </span>
          </h1>
        </Transition>

        <Transition name="reveal-up" appear>
          <p class="text-xl md:text-2xl text-slate-300 max-w-3xl mx-auto font-light leading-relaxed mb-10">
            Delivering premium quality cellulose fiber to global standards through
            continuous innovation and high environmental responsibility.
          </p>
        </Transition>

        <Transition name="reveal-up" appear>
          <button
            @click="scrollToProducts"
            class="px-10 py-5 bg-emerald-600 hover:bg-emerald-500 text-white rounded-full text-lg font-bold transition-all transform hover:scale-105 hover:shadow-[0_0_20px_rgba(16,185,129,0.4)] flex items-center gap-2 mx-auto cursor-pointer border-none"
          >
            Explore Products
            <svg class="w-6 h-6 animate-bounce-slow" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7-7v18"/>
            </svg>
          </button>
        </Transition>
      </div>
    </header>

    <!-- ─── PRODUCTS GRID ─────────────────────────────────────── -->
    <main id="product-target" class="container mx-auto px-6 lg:px-16 mt-4 relative z-20 pb-24 scroll-mt-10">

      <!-- Loading skeleton -->
      <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        <div v-for="i in 4" :key="i" class="rounded-[2rem] bg-white shadow-xl h-[440px] animate-pulse">
          <div class="h-64 bg-slate-200 rounded-t-[2rem]"></div>
          <div class="p-8 space-y-3">
            <div class="h-5 bg-slate-200 rounded w-3/4"></div>
            <div class="h-4 bg-slate-100 rounded w-full"></div>
            <div class="h-4 bg-slate-100 rounded w-5/6"></div>
          </div>
        </div>
      </div>

      <!-- Empty -->
      <div v-else-if="products.length === 0" class="text-center py-24 text-slate-400">
        <p class="text-5xl mb-4">📦</p>
        <p class="text-xl">No products available yet.</p>
      </div>

      <!-- Grid -->
      <TransitionGroup v-else tag="div" name="list-fade"
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        <article
          v-for="(product, index) in products" :key="product.ID || index"
          class="group bg-white rounded-[2rem] shadow-xl shadow-slate-200/50 border border-slate-100 overflow-hidden transition-all duration-500 hover:-translate-y-4 flex flex-col h-full">

          <!-- Image -->
          <div class="relative h-64 overflow-hidden bg-slate-100">
            <img
              :src="getImageUrl(product.thumbnail_path || product.image_url)"
              :alt="product.name || product.title"
              class="w-full h-full object-cover transition-transform duration-1000 group-hover:scale-110"
              @error="(e) => e.target.src = fallbackImg"
            />
            <div class="absolute inset-0 bg-black/10 group-hover:bg-transparent transition-all duration-500"></div>

            <!-- Category badge -->
            <div class="absolute top-4 right-4 z-10">
              <div class="px-4 py-1.5 rounded-full bg-white/90 backdrop-blur-md text-xs font-bold text-emerald-800 uppercase tracking-wider shadow-sm">
                {{ product.category || 'Pulp' }}
              </div>
            </div>

            <!-- Icon -->
            <div class="absolute bottom-6 left-6 z-10 p-3 rounded-2xl bg-emerald-600 text-white shadow-lg transform -rotate-6 group-hover:rotate-0 transition-all duration-500">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" :d="categoryIcon(product.category)"/>
              </svg>
            </div>
          </div>

          <!-- Body -->
          <div class="p-8 flex flex-col flex-grow">
            <h3 class="text-xl font-bold text-slate-900 mb-3 group-hover:text-emerald-600 transition-colors leading-tight">
              {{ product.name || product.title }}
            </h3>

            <p class="text-sm text-slate-600 font-normal leading-relaxed mb-5 flex-grow line-clamp-3">
              {{ product.summary || product.description }}
            </p>

            <!-- Tags -->
            <div class="flex flex-wrap gap-2 mb-6" v-if="parseTags(product.tags).length">
              <span v-for="tag in parseTags(product.tags)" :key="tag"
                class="text-xs font-semibold bg-slate-100 text-slate-600 px-3 py-1 rounded-md">
                {{ tag }}
              </span>
            </div>

            <!-- Detail link -->
            <RouterLink
              :to="`/product/${product.ID || product.id}`"
              class="inline-flex items-center justify-between text-sm font-bold text-emerald-700 hover:text-emerald-500 group/btn pt-4 border-t border-slate-100 w-full no-underline">
              <span>Detail Produk</span>
              <div class="w-9 h-9 rounded-full border border-emerald-100 flex items-center justify-center bg-emerald-50 text-emerald-700 group-hover/btn:bg-emerald-600 group-hover/btn:text-white transition-all duration-300">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
                </svg>
              </div>
            </RouterLink>
          </div>
        </article>
      </TransitionGroup>
    </main>

    <!-- ─── SLIDER / FACILITY ─────────────────────────────────── -->
    <section class="container mx-auto px-6 lg:px-16 pb-24">
      <div class="relative rounded-[3rem] overflow-hidden shadow-2xl h-[520px] group bg-[#14532d]">
        <TransitionGroup name="fade">
          <div v-for="(slide, i) in slides" :key="i" v-show="currentSlide === i" class="absolute inset-0">
            <img :src="slide.url" class="w-full h-full object-contain opacity-80 p-4"
              @error="(e) => e.target.src='https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?q=80&w=1200'" />
            <div class="absolute inset-0 bg-gradient-to-r from-black/80 via-black/20 to-transparent"></div>
            <div class="absolute inset-y-0 left-0 w-full md:w-2/3 flex items-center p-12 md:p-24 z-20">
              <div class="max-w-xl">
                <Transition name="reveal-up">
                  <div v-if="currentSlide === i" class="space-y-5">
                    <span class="text-emerald-400 font-bold tracking-[0.3em] text-sm uppercase">Industrial Excellence</span>
                    <h2 class="text-4xl md:text-5xl font-serif font-bold text-white leading-tight">{{ slide.title }}</h2>
                    <p class="text-slate-300 text-lg font-light leading-relaxed">{{ slide.desc }}</p>
                  </div>
                </Transition>
              </div>
            </div>
          </div>
        </TransitionGroup>

        <!-- Progress bar -->
        <div class="absolute bottom-0 left-0 h-1 bg-emerald-500 z-40 transition-none" :style="{ width: progress + '%' }"></div>

        <!-- Arrows -->
        <div class="absolute bottom-10 right-10 z-30 flex gap-3 opacity-0 group-hover:opacity-100 transition-opacity duration-500">
          <button @click="manualNav(prevSlide)" class="p-3.5 rounded-full bg-white/10 backdrop-blur-md text-white border border-white/20 hover:bg-emerald-600 transition-all cursor-pointer">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/></svg>
          </button>
          <button @click="manualNav(nextSlide)" class="p-3.5 rounded-full bg-white/10 backdrop-blur-md text-white border border-white/20 hover:bg-emerald-600 transition-all cursor-pointer">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/></svg>
          </button>
        </div>
 
        <!-- Dots -->
        <div class="absolute bottom-10 left-12 md:left-24 z-30 flex gap-3">
          <button v-for="(_, i) in slides" :key="i" @click="manualNav(() => goToSlide(i))"
            class="h-1 transition-all duration-500 rounded-full cursor-pointer"
            :class="currentSlide === i ? 'w-12 bg-emerald-500' : 'w-4 bg-white/20 hover:bg-white/40'">
          </button>
        </div>
      </div>
    </section>

  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { RouterLink } from 'vue-router'
import api from '../services/api'

const BASE_URL = 'http://localhost:8080'
const fallbackImg = 'https://images.unsplash.com/photo-1605281317010-fe5ffe798156?q=80&w=600'

const products = ref([])
const loading  = ref(true)

// ─── Fetch products ───────────────────────────────────────────
onMounted(async () => {
  try {
    const res = await api.get('/products')
    products.value = res.data?.data || res.data || []
  } catch {
    products.value = []
  } finally {
    loading.value = false
  }
  startTimer()
})
onUnmounted(() => stopTimer())

const getImageUrl = (path) => {
  if (!path) return fallbackImg
  if (path.startsWith('http')) return path
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const parseTags = (tags) => {
  if (!tags) return []
  if (Array.isArray(tags)) return tags
  try { return JSON.parse(tags) } catch { return tags.split(',').map(t => t.trim()) }
}

const categoryIcon = (cat) => {
  const c = (cat || '').toLowerCase()
  if (c.includes('pulp'))    return 'M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10'
  if (c.includes('tissue'))  return 'M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z'
  if (c.includes('paper'))   return 'M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253'
  return 'M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z'
}

const scrollToProducts = () => {
  document.getElementById('product-target')?.scrollIntoView({ behavior: 'smooth', block: 'start' })
}

// ─── Slider ───────────────────────────────────────────────────
const currentSlide = ref(0)
const progress = ref(0)
let timer = null, progressTimer = null

const slides = [
  { url: '/images/produk.png', title: 'Smart Warehousing System', desc: 'Automated warehouse management ensuring timely and safe logistics availability.' },
  { url: '/images/produk2.jpeg', title: 'Sustainable Production', desc: 'Modern production processes with low carbon footprint to preserve future forest ecosystems.' },
  { url: '/images/produk3.jpeg', title: 'Global Quality Control', desc: 'State-of-the-art laboratories monitoring every fiber micron for our industry partners.' },
  { url: '/images/pulp.jpeg', title: 'Premium Pulp Production', desc: 'High-quality cellulose fiber produced through sustainable and eco-friendly processes.' },
]

const startTimer = () => {
  timer = setInterval(nextSlide, 6000)
  const step = 100 / (6000 / 10)
  progressTimer = setInterval(() => { progress.value = Math.min(progress.value + step, 100) }, 10)
}
const stopTimer  = () => { clearInterval(timer); clearInterval(progressTimer) }
const nextSlide  = () => { currentSlide.value = (currentSlide.value + 1) % slides.length; progress.value = 0 }
const prevSlide  = () => { currentSlide.value = (currentSlide.value - 1 + slides.length) % slides.length; progress.value = 0 }
const goToSlide  = (i) => { currentSlide.value = i; progress.value = 0 }
const manualNav  = (fn) => { stopTimer(); fn(); startTimer() }
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@300;400;600;700&family=Playfair+Display:wght@700&display=swap');
.font-serif { font-family: 'Playfair Display', serif; }

.list-fade-enter-active { transition: all 0.6s cubic-bezier(0.34,1.56,0.64,1); }
.list-fade-enter-from   { opacity: 0; transform: scale(0.85) translateY(30px); }
.reveal-up-enter-active   { transition: all 1s cubic-bezier(0.22,1,0.36,1); }
.reveal-up-enter-from     { opacity: 0; transform: translateY(40px); }
.reveal-down-enter-active { transition: all 1s cubic-bezier(0.22,1,0.36,1); }
.reveal-down-enter-from   { opacity: 0; transform: translateY(-40px); }
.fade-enter-active, .fade-leave-active { transition: opacity 1.2s ease-in-out; }
.fade-enter-from, .fade-leave-to       { opacity: 0; }

@keyframes subtle-zoom { 0% { transform:scale(1); } 100% { transform:scale(1.08); } }
.animate-subtle-zoom { animation: subtle-zoom 20s infinite alternate ease-in-out; }
@keyframes bounce { 0%,100%{transform:translateY(-5%);} 50%{transform:none;} }
.animate-bounce-slow { animation: bounce 2s infinite; }
</style>