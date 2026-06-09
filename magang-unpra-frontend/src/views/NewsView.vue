<template>
  <section class="pt-24 pb-16 bg-gray-900 text-white">
    <div class="max-w-7xl mx-auto px-10">
      <p class="text-orange-400 text-xs font-semibold tracking-widest mb-2">STRATEGIC MEDIA</p>
      <h1 class="text-5xl font-black">Latest News</h1>
      <p class="text-gray-400 mt-3">Inside look at our latest reforestation projects, industrial innovations, and community impact.</p>
    </div>
  </section>

  <section class="py-24 bg-white">
    <div class="max-w-7xl mx-auto px-10 space-y-10">

      <!-- Loading -->
      <div v-if="loading" class="text-center text-gray-400 py-20">Loading...</div>

      <!-- Empty -->
      <div v-else-if="newsList.length === 0" class="text-center text-gray-400 py-20">
        Belum ada berita.
      </div>

      <!-- News List -->
      <div v-else v-for="news in newsList" :key="news.ID" class="flex gap-8 border-b border-gray-100 pb-10">
        <div class="w-64 h-48 bg-gray-200 rounded-2xl flex-shrink-0 flex items-center justify-center">
          <p class="text-gray-400 text-sm">Image</p>
        </div>
        <div class="flex-1">
          <p class="text-green-600 text-xs font-semibold tracking-widest mb-2">{{ news.category }}</p>
          <h3 class="text-2xl font-bold text-gray-900 mb-3">{{ news.title }}</h3>
          <p class="text-gray-600 leading-relaxed mb-6">{{ news.content }}</p>
          <button class="border border-gray-900 px-6 py-2 rounded-full text-sm font-medium hover:bg-gray-900 hover:text-white transition">
            Continue Reading →
          </button>
        </div>
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