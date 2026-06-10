<template>
  <!-- Hero -->
  <section class="relative h-64 flex items-center justify-center bg-cover bg-center"
    style="background-image: url('https://images.unsplash.com/photo-1448375240586-882707db888b?w=1920')">
    <div class="absolute inset-0 bg-black/60"></div>
    <div class="relative z-10 text-center text-white">
      <p class="text-sm text-gray-300 mb-2">
        You are here: <span class="text-white font-medium">Home » News</span>
      </p>
      <h1 class="text-4xl font-black tracking-wide">
        CATEGORY: <span class="text-green-400">NEWS</span>
      </h1>
    </div>
  </section>

  <!-- News List -->
  <section class="py-16 bg-white">
    <div class="max-w-4xl mx-auto px-6">

      <!-- Loading -->
      <div v-if="loading" class="text-center py-20">
        <div class="w-8 h-8 border-4 border-green-600 border-t-transparent rounded-full animate-spin mx-auto"></div>
        <p class="text-gray-400 mt-4 text-sm">Loading berita...</p>
      </div>

      <!-- Empty -->
      <div v-else-if="newsList.length === 0" class="text-center py-20">
        <p class="text-5xl mb-4">📰</p>
        <p class="text-gray-500">Belum ada berita.</p>
      </div>

      <!-- News Items -->
      <div v-else class="space-y-10">
        <article v-for="news in newsList" :key="news.ID"
          class="flex gap-8 pb-10 border-b border-gray-100 last:border-0">

          <!-- Thumbnail -->
          <div class="w-56 h-40 flex-shrink-0 rounded-xl overflow-hidden bg-gray-100">
            <img
              v-if="news.Images && news.Images.length > 0"
              :src="news.Images[0].image_url"
              :alt="news.title"
              class="w-full h-full object-cover"
            />
            <div v-else
              class="w-full h-full bg-gradient-to-br from-green-50 to-green-100 flex items-center justify-center">
              <span class="text-5xl">🌿</span>
            </div>
          </div>

          <!-- Content -->
          <div class="flex-1">
            <div class="flex items-center gap-3 mb-2 text-xs text-gray-500">
              <span>👤 Posted by Admin</span>
              <span>•</span>
              <span class="bg-green-600 text-white px-2 py-0.5 rounded font-semibold">
                {{ news.category || 'News' }}
              </span>
            </div>
            <h2 class="text-xl font-bold text-gray-900 mb-3 leading-snug hover:text-green-600 transition cursor-pointer">
              {{ news.title }}
            </h2>
            <p class="text-gray-500 text-sm leading-relaxed mb-5 line-clamp-3">
              {{ news.content }}
            </p>
            <button
              class="border border-gray-400 text-gray-600 px-5 py-2 text-xs font-semibold tracking-widest uppercase hover:bg-gray-900 hover:text-white hover:border-gray-900 transition">
              Continue Reading
            </button>
          </div>
        </article>
      </div>

    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const newsList = ref([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await api.get('/news')
    newsList.value = res.data.data
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
})
</script>