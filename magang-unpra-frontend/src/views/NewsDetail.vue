<template>
  <!-- Hero -->
  <section class="relative h-44 md:h-64 flex items-center justify-center bg-cover bg-center"
    style="background-image: url('https://images.unsplash.com/photo-1448375240586-882707db888b?w=1920')">
    <div class="absolute inset-0 bg-black/60"></div>
    <div class="relative z-10 text-center text-white px-4">
      <p class="text-xs sm:text-sm text-gray-300 mb-2">
        You are here:
        <span class="cursor-pointer hover:underline" @click="router.push('/news')">Home » News</span>
        » <span class="text-white font-medium hidden sm:inline">{{ news?.title }}</span>
      </p>
      <h1 class="text-xl md:text-3xl font-black tracking-wide">Detail Berita</h1>
    </div>
  </section>

  <!-- Loading -->
  <div v-if="loading" class="text-center py-20">
    <div class="w-8 h-8 border-4 border-[#5F9E42] border-t-transparent rounded-full animate-spin mx-auto"></div>
    <p class="text-gray-400 mt-4 text-sm">Memuat berita...</p>
  </div>

  <!-- Konten -->
  <section v-else-if="news" class="py-8 md:py-16 bg-white">
    <div class="max-w-3xl mx-auto px-4 md:px-6">

      <!-- Meta -->
      <div class="flex flex-wrap items-center gap-2 sm:gap-3 mb-4 text-xs text-gray-500">
        <span>Posted by Admin</span>
        <span>•</span>
        <span>{{ formatDetailDate(news.published_at || news.PublishedAt || news.CreatedAt || news.created_at) }}</span>
        <span>•</span>
        <span class="bg-[#5F9E42] text-white px-2 py-0.5 rounded font-semibold">
          {{ news.category || 'News' }}
        </span>
      </div>

      <!-- Judul -->
      <h1 class="text-2xl md:text-3xl font-black text-gray-900 mb-4 md:mb-6 leading-snug">
        {{ news.title }}
      </h1>

      <!-- Gambar utama -->
      <div v-if="news.Images && news.Images.length > 0"
        class="w-full rounded-xl overflow-hidden mb-6 md:mb-8 max-h-64 md:max-h-96">
        <img :src="getImageUrl(news.Images[0].image_url)" :alt="news.title"
          class="w-full h-full object-cover"/>
      </div>
      <div v-else-if="news.thumbnail_path"
        class="w-full rounded-xl overflow-hidden mb-6 md:mb-8 max-h-64 md:max-h-96">
        <img :src="getImageUrl(news.thumbnail_path)" :alt="news.title"
          class="w-full h-full object-cover"/>
      </div>

      <!-- Isi konten -->
      <div class="prose prose-gray max-w-none text-gray-700 leading-relaxed text-[14px] md:text-[15px]"
        v-html="news.content">
      </div>

      <!-- Galeri foto tambahan -->
      <div v-if="news.Images && news.Images.length > 1"
        class="grid grid-cols-1 sm:grid-cols-2 gap-3 sm:gap-4 mt-6 md:mt-10">
        <div v-for="(img, i) in news.Images.slice(1)" :key="i"
          class="rounded-lg overflow-hidden h-56 sm:h-48">
          <img :src="getImageUrl(img.image_url)" class="w-full h-full object-cover"/>
        </div>
      </div>

      <!-- Tombol kembali -->
      <div class="mt-8 md:mt-12">
        <button @click="router.push('/news')"
          class="border border-gray-400 text-gray-600 px-4 md:px-5 py-2 text-xs font-semibold tracking-widest uppercase hover:bg-gray-900 hover:text-white hover:border-gray-900 transition">
          ← Kembali ke News
        </button>
      </div>

    </div>
  </section>

  <!-- Not found -->
  <div v-else class="text-center py-20">
    <p class="text-5xl mb-4">😕</p>
    <p class="text-gray-500">Berita tidak ditemukan.</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import api from '../services/api'

const BASE_URL = 'http://localhost:8080'
const route  = useRoute()
const router = useRouter()
const news   = ref(null)
const loading = ref(true)

const getImageUrl = (path) => {
  if (!path) {return ''}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const formatDetailDate = (date) => {
  if (!date) {return ''}
  return new Date(date).toLocaleDateString('id-ID', { day: '2-digit', month: 'short', year: 'numeric' })
}

onMounted(async () => {
  try {
    const res = await api.get(`/news/${route.params.id}`)
    news.value = res.data.data
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
})
</script>