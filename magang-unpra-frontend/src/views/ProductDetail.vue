<template>
  <div class="min-h-screen bg-slate-50">

    <!-- Loading -->
    <div v-if="loading" class="flex items-center justify-center min-h-screen">
      <div class="w-10 h-10 border-4 border-slate-200 border-t-emerald-500 rounded-full animate-spin"></div>
    </div>

    <!-- Not found -->
    <div v-else-if="!product" class="flex flex-col items-center justify-center min-h-screen text-slate-400 gap-4">
      <p class="text-6xl">📦</p>
      <p class="text-xl font-semibold">Product not found.</p>
      <RouterLink to="/product" class="text-emerald-600 font-bold hover:underline">← Back to Products</RouterLink>
    </div>

    <!-- Content -->
    <div v-else>
      <!-- ─── Hero ── -->
      <PageHero
        :title="product.name"
        subtitle="Product detail and specifications."
        :breadcrumbs="[{ label: 'Home', to: '/' }, { label: 'Product', to: '/product' }, { label: product.name }]"
      />

      <!-- ─── Detail Content ── -->
      <main class="container mx-auto px-6 lg:px-16 py-16 -mt-4 relative z-10">
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-12">

          <!-- Left: main info -->
          <div class="lg:col-span-2 space-y-10">

            <!-- Summary -->
            <div class="bg-white rounded-3xl p-8 shadow-sm border border-slate-100">
              <h2 class="text-xl font-bold text-slate-900 mb-4 flex items-center gap-2">
                <span class="w-1.5 h-6 bg-emerald-500 rounded-full block"></span>
                Product Overview
              </h2>
              <p class="product-detail-content" v-html="product.summary || product.description || 'No description available.'">
              </p>
            </div>

            <!-- Full content -->
            <div v-if="product.description && product.summary" class="bg-white rounded-3xl p-8 shadow-sm border border-slate-100">
              <h2 class="text-xl font-bold text-slate-900 mb-4 flex items-center gap-2">
                <span class="w-1.5 h-6 bg-emerald-500 rounded-full block"></span>
                Detailed Description
              </h2>
              <div class="product-detail-content" v-html="product.description">
              </div>
            </div>

            <!-- Images gallery -->
            <div v-if="product.Images && product.Images.length" class="bg-white rounded-3xl p-8 shadow-sm border border-slate-100">
              <h2 class="text-xl font-bold text-slate-900 mb-6 flex items-center gap-2">
                <span class="w-1.5 h-6 bg-emerald-500 rounded-full block"></span>
                Product Gallery
              </h2>
              <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
                <div v-for="(img, i) in product.Images" :key="i"
                  class="aspect-square rounded-2xl overflow-hidden bg-slate-100 cursor-pointer hover:scale-[1.02] transition-transform"
                  @click="lightboxImg = getImageUrl(img.image_url); showLightbox = true">
                  <img :src="getImageUrl(img.image_url)" :alt="`Image ${i+1}`"
                    class="w-full h-full object-cover"
                    @error="(e) => e.target.src = fallbackImg" />
                </div>
              </div>
            </div>
          </div>

          <!-- Right: specs sidebar -->
          <div class="space-y-6">

            <!-- Tags / Specs -->
            <div class="bg-white rounded-3xl p-6 shadow-sm border border-slate-100">
              <h3 class="text-base font-bold text-slate-900 mb-4 flex items-center gap-2">
                <span class="w-1.5 h-5 bg-emerald-500 rounded-full block"></span>
                Technical Specifications
              </h3>
              <div class="flex flex-wrap gap-2">
                <span v-for="tag in parseTags(product.tags)" :key="tag"
                  class="text-xs font-bold bg-emerald-50 text-emerald-700 px-3.5 py-1.5 rounded-lg border border-emerald-100">
                  {{ tag }}
                </span>
                <p v-if="!parseTags(product.tags).length" class="text-slate-400 text-sm">
                  No specifications listed.
                </p>
              </div>
            </div>

            <!-- Info card -->
            <div class="bg-emerald-600 rounded-3xl p-6 text-white">
              <div class="w-10 h-10 bg-white/20 rounded-xl flex items-center justify-center mb-4">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
              </div>
              <h4 class="font-bold text-lg mb-2">Need more info?</h4>
              <p class="text-emerald-100 text-sm leading-relaxed mb-5">
                Contact our team for technical specifications, pricing, and custom requirements.
              </p>
              <RouterLink to="/contact"
                class="block text-center bg-white text-emerald-700 font-bold py-3 px-4 rounded-xl text-sm hover:bg-emerald-50 transition-colors no-underline">
                Contact Us →
              </RouterLink>
            </div>

            <!-- Category -->
            <div class="bg-white rounded-3xl p-6 shadow-sm border border-slate-100">
              <h3 class="text-sm font-bold text-slate-500 uppercase tracking-widest mb-3">Category</h3>
              <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-emerald-50 rounded-xl flex items-center justify-center">
                  <svg class="w-5 h-5 text-emerald-600" fill="none" stroke="currentColor" stroke-width="1.8" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z"/>
                  </svg>
                </div>
                <span class="font-bold text-slate-800">{{ product.category || 'General' }}</span>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>

    <!-- ─── Lightbox ── -->
    <Transition name="modal-fade">
      <div v-if="showLightbox"
        class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 backdrop-blur-sm p-4"
        @click="showLightbox = false">
        <img :src="lightboxImg" class="max-w-full max-h-[90vh] rounded-2xl shadow-2xl" @click.stop />
        <button @click="showLightbox = false"
          class="absolute top-6 right-6 text-white bg-white/10 hover:bg-white/20 border border-white/20 rounded-full p-2 cursor-pointer">
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
          </svg>
        </button>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import PageHero from '../components/PageHero.vue'
import { ref, onMounted } from 'vue'
import { useRoute, RouterLink } from 'vue-router'
import api from '../services/api'

const BASE_URL = import.meta.env.VITE_BASE_URL || 'http://localhost:8080'
const fallbackImg = 'https://images.unsplash.com/photo-1605281317010-fe5ffe798156?q=80&w=600'

const route   = useRoute()
const product = ref(null)
const loading = ref(true)
const showLightbox = ref(false)
const lightboxImg  = ref('')

onMounted(async () => {
  try {
    const res = await api.get(`/products/${route.params.id}`)
    product.value = res.data?.data || res.data
  } catch {
    product.value = null
  } finally {
    loading.value = false
  }
})

const getImageUrl = (path) => {
  if (!path) {return fallbackImg}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const parseTags = (tags) => {
  if (!tags) {return []}
  if (Array.isArray(tags)) {return tags}
  try { return JSON.parse(tags) } catch { return String(tags).split(',').map(t => t.trim()) }
}
</script>

<style scoped>
@keyframes subtle-zoom { 0%{transform:scale(1);}100%{transform:scale(1.06);} }
.animate-subtle-zoom { animation: subtle-zoom 12s infinite alternate ease-in-out; }
.modal-fade-enter-active,.modal-fade-leave-active{transition:opacity 0.3s ease;}
.modal-fade-enter-from,.modal-fade-leave-to{opacity:0;}
</style>

<style>
.product-detail-content { color: #475569; font-size: 15px; line-height: 1.75; }
.product-detail-content p { margin-bottom: 1em; }
.product-detail-content h1 { font-size: 1.5em; font-weight: 700; margin: 0.8em 0 0.4em; color: #111827; }
.product-detail-content h2 { font-size: 1.25em; font-weight: 600; margin: 0.8em 0 0.4em; color: #111827; }
.product-detail-content h3 { font-size: 1.1em; font-weight: 600; margin: 0.8em 0 0.4em; color: #111827; }
.product-detail-content ul { list-style-type: disc; padding-left: 1.5em; margin: 0.5em 0; }
.product-detail-content ol { list-style-type: decimal; padding-left: 1.5em; margin: 0.5em 0; }
.product-detail-content li { margin: 0.25em 0; }
.product-detail-content img { max-width: 100%; border-radius: 8px; margin: 1em 0; }
.product-detail-content blockquote { border-left: 3px solid #10b981; padding-left: 12px; margin: 0.75em 0; color: #6b7280; }
.product-detail-content strong { font-weight: 700; color: #111827; }
</style>