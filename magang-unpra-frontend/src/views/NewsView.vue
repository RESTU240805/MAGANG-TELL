<template>
  <div class="min-h-screen bg-white font-sans">

    <!-- ─── HERO ──────────────────────────────────────────────── -->
    <section class="relative h-[400px] bg-[#1a3a2a] overflow-hidden">
      <img
        src="/images/gedung.jpeg"
        alt="News Hero"
        class="absolute inset-0 w-full h-full object-cover opacity-50 animate-hero-zoom"
        @error="(e) => e.target.src='https://images.unsplash.com/photo-1541888946425-d81bb19240f5?q=80&w=1600'"
      />
      <div class="absolute inset-0 bg-gradient-to-r from-black/70 via-black/30 to-transparent"></div>
      <div class="absolute inset-0 flex items-end">
        <div class="max-w-[1200px] mx-auto px-8 pb-12 w-full">
          <p class="text-white/70 text-sm mb-2 opacity-0 animate-fade-up" style="animation-delay: 0.1s">
            <RouterLink to="/" class="hover:text-white no-underline text-white/70">Home</RouterLink>
            <span class="mx-2">/</span>
            <span class="text-white">News</span>
          </p>
          <h1 class="text-4xl md:text-5xl font-black text-white mb-4 opacity-0 animate-fade-up" style="animation-delay: 0.2s">News &amp; Update</h1>
          <p class="text-white/80 text-base max-w-lg opacity-0 animate-fade-up" style="animation-delay: 0.3s">
            Berita terbaru dan informasi resmi<br>
            PT Tanjungenim Lestari Pulp and Paper.
          </p>
        </div>
      </div>
    </section>

    <!-- ─── MAIN CONTENT ──────────────────────────────────────── -->
    <div class="max-w-[1200px] mx-auto px-8 py-16">

      <!-- Loading -->
      <div v-if="loading" class="flex justify-center py-20">
        <div class="w-10 h-10 border-4 border-gray-200 border-t-[#5F9E42] rounded-full animate-spin"></div>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="text-center py-20 text-gray-400">
        <p class="text-xl mb-4">Gagal memuat berita.</p>
        <button @click="fetchNews" class="px-6 py-2 bg-[#5F9E42] text-white rounded text-sm font-semibold cursor-pointer">Coba Lagi</button>
      </div>

      <!-- Empty -->
      <div v-else-if="!newsList.length" class="text-center py-20 text-gray-400">
        <p class="text-4xl mb-4">📰</p>
        <p class="text-xl">Belum ada berita.</p>
      </div>

      <!-- Content -->
      <div v-else class="flex gap-10">

        <!-- Left: Articles -->
        <div class="flex-1 min-w-0">
          <h2 class="text-2xl font-bold text-gray-900 mb-2 anim-item">Berita Terbaru</h2>
          <div class="w-16 h-1 bg-[#5F9E42] rounded mb-8 anim-item"></div>

          <!-- News Items -->
          <div v-for="(item, index) in newsList" :key="item.id"
            class="flex gap-6 pb-8 mb-8 border-b border-gray-200 last:border-b-0 last:mb-0 last:pb-0 anim-item"
            :style="{ animationDelay: (index * 0.1) + 's' }">

            <!-- Image -->
            <div class="w-[340px] h-[220px] flex-shrink-0 rounded-lg overflow-hidden bg-gray-100">
              <img :src="getImageUrl(item.image)" :alt="item.title"
                class="w-full h-full object-cover"
                @error="(e) => e.target.src = fallbackImg" />
            </div>

            <!-- Content -->
            <div class="flex flex-col justify-start pt-1 min-w-0">
              <span class="inline-block text-[10px] font-bold tracking-widest uppercase text-white bg-[#5F9E42] px-3 py-1 rounded mb-3 w-fit">
                {{ item.category || 'News' }}
              </span>
              <div class="flex items-center gap-4 text-xs text-gray-500 mb-3">
                <span class="flex items-center gap-1">
                  <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                    <path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2"/>
                    <circle cx="12" cy="7" r="4"/>
                  </svg>
                  {{ item.author || 'Admin' }}
                </span>
                <span class="flex items-center gap-1">
                  <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
                    <rect x="3" y="4" width="18" height="18" rx="2"/>
                    <line x1="16" y1="2" x2="16" y2="6"/><line x1="8" y1="2" x2="8" y2="6"/>
                    <line x1="3" y1="10" x2="21" y2="10"/>
                  </svg>
                  {{ formatDate(item.date) }}
                </span>
              </div>
              <h3 class="text-lg font-bold text-gray-900 mb-2 leading-snug">{{ item.title }}</h3>
              <p class="text-sm text-gray-600 leading-relaxed mb-3 line-clamp-3">{{ item.excerpt }}</p>
              <RouterLink :to="`/news/${item.id}`"
                class="text-[#5F9E42] text-sm font-bold no-underline hover:underline">
                Baca Selengkapnya →
              </RouterLink>
            </div>
          </div>

          <!-- Pagination -->
          <div v-if="totalPages > 1" class="flex items-center gap-2 mt-10 justify-center">
            <button @click="prevPage" :disabled="currentPage <= 1"
              class="w-9 h-9 rounded border border-gray-300 bg-white text-gray-600 text-sm cursor-pointer hover:bg-gray-100 disabled:opacity-30 disabled:cursor-not-allowed flex items-center justify-center">
              ‹
            </button>
            <button v-for="page in totalPages" :key="page" @click="goToPage(page)"
              class="w-9 h-9 rounded text-sm font-semibold cursor-pointer flex items-center justify-center transition-colors"
              :class="currentPage === page
                ? 'bg-[#5F9E42] text-white border border-[#5F9E42]'
                : 'bg-white text-gray-600 border border-gray-300 hover:bg-gray-100'">
              {{ page }}
            </button>
            <button @click="nextPage" :disabled="currentPage >= totalPages"
              class="w-9 h-9 rounded border border-gray-300 bg-white text-gray-600 text-sm cursor-pointer hover:bg-gray-100 disabled:opacity-30 disabled:cursor-not-allowed flex items-center justify-center">
              ›
            </button>
          </div>
        </div>

        <!-- Right: Sidebar -->
        <aside class="w-[320px] flex-shrink-0 anim-slide-right">

          <!-- Search -->
          <div class="relative mb-8">
            <input v-model="searchQuery" type="text" placeholder="Cari berita..."
              class="w-full px-4 py-3 pr-12 border border-gray-300 rounded-lg text-sm focus:outline-none focus:border-[#5F9E42]" />
            <button class="absolute right-0 top-0 h-full w-12 bg-[#5F9E42] text-white rounded-r-lg flex items-center justify-center cursor-pointer hover:bg-[#4d8536]">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24">
                <circle cx="11" cy="11" r="8"/><line x1="21" y1="21" x2="16.65" y2="16.65"/>
              </svg>
            </button>
          </div>

          <!-- Berita Terbaru -->
          <div>
            <h3 class="text-lg font-bold text-gray-900 mb-5">Berita Terbaru</h3>
            <div v-for="item in newsList.slice(0, 5)" :key="'side-' + item.id"
              class="flex gap-3 pb-4 mb-4 border-b border-gray-100 last:border-b-0 last:mb-0 last:pb-0">
              <div class="w-[90px] h-[65px] flex-shrink-0 rounded overflow-hidden bg-gray-100">
                <img :src="getImageUrl(item.image)" :alt="item.title"
                  class="w-full h-full object-cover"
                  @error="(e) => e.target.src = fallbackImg" />
              </div>
              <div class="flex flex-col justify-center min-w-0">
                <h4 class="text-sm font-bold text-gray-900 leading-snug mb-1 line-clamp-2">{{ item.title }}</h4>
                <span class="text-[11px] text-gray-400">{{ formatDate(item.date) }}</span>
              </div>
            </div>
          </div>

        </aside>
      </div>
    </div>

    <footer class="site-footer">
      <div class="footer-container">
        <div class="footer-left-content"></div>
        <div class="footer-copyright">
          <p>Copyright 2026 PT TELPP. All right reserved.</p>
        </div>
      </div>
    </footer>

  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { RouterLink } from 'vue-router'
import api from '../services/api'

const BASE_URL = 'http://localhost:8080'
const fallbackImg = 'https://placehold.co/600x400/e8e8e8/999?text=News'

const newsList = ref([])
const loading = ref(true)
const error = ref(false)
const currentPage = ref(1)
const totalPages = ref(1)
const searchQuery = ref('')

const formatDate = (dateStr) => {
  if (!dateStr) {return ''}
  const date = new Date(dateStr)
  return date.toLocaleDateString('en-GB', { day: '2-digit', month: 'short', year: 'numeric' })
}

const getImageUrl = (path) => {
  if (!path) {return fallbackImg}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const fetchNews = async () => {
  loading.value = true
  error.value = false
  try {
    const res = await api.get(`/news?page=${currentPage.value}`)
    const data = res.data

    const rawList = data.data || data.news || data || []
    newsList.value = rawList
      .filter(item => item.is_published)
      .map(item => ({
        id: item.ID || item.id,
        title: item.title,
        excerpt: item.summary || '',
        category: item.category,
        author: item.author || 'Admin',
        date: item.CreatedAt || item.created_at || item.date,
        image: item.Images?.[0]?.image_url || item.thumbnail_path || '',
      }))

    if (data.total_pages) {totalPages.value = data.total_pages}
  } catch (e) {
    console.error(e)
    error.value = true
  } finally {
    loading.value = false
  }
}

const goToPage = (page) => {
  currentPage.value = page
  fetchNews()
}

const prevPage = () => {
  if (currentPage.value > 1) {goToPage(currentPage.value - 1)}
}

const nextPage = () => {
  if (currentPage.value < totalPages.value) {goToPage(currentPage.value + 1)}
}

let observer = null

onMounted(async () => {
  await fetchNews()
  await nextTick()
  initObserver()
})

onUnmounted(() => {
  if (observer) {observer.disconnect()}
})

const initObserver = () => {
  observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.add('anim-visible')
        observer.unobserve(entry.target)
      }
    })
  }, { threshold: 0.15 })

  document.querySelectorAll('.anim-item, .anim-slide-right').forEach(el => {
    observer.observe(el)
  })
}
</script>

<style scoped>
.animate-spin {
  animation: spin 0.8s linear infinite;
}
@keyframes spin {
  to { transform: rotate(360deg); }
}

/* Hero animations */
@keyframes heroZoom {
  0% { transform: scale(1); }
  100% { transform: scale(1.08); }
}
.animate-hero-zoom {
  animation: heroZoom 20s infinite alternate ease-in-out;
}

@keyframes fadeUp {
  from { opacity: 0; transform: translateY(24px); }
  to { opacity: 1; transform: translateY(0); }
}
.animate-fade-up {
  animation: fadeUp 0.8s cubic-bezier(0.22, 1, 0.36, 1) forwards;
}

/* Scroll animations */
.anim-item {
  opacity: 0;
  transform: translateY(30px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-item.anim-visible {
  opacity: 1;
  transform: translateY(0);
}

.anim-slide-right {
  opacity: 0;
  transform: translateX(40px);
  transition: opacity 0.7s cubic-bezier(0.22, 1, 0.36, 1), transform 0.7s cubic-bezier(0.22, 1, 0.36, 1);
}
.anim-slide-right.anim-visible {
  opacity: 1;
  transform: translateX(0);
}
/* ============ FOOTER ============ */
.site-footer {
  width: 100%;
  /* Warna hijau tua disesuaikan dengan identitas PT TEL pada gambar */
  background-color: #5F9E42; 
  padding: 25px 0; /* Memberikan ruang tinggi baris yang pas */
  margin-top: 40px; /* Jarak pemisah dari konten di atasnya */
}

.footer-container {
  display: flex;
  align-items: center;
  justify-content: center; /* DIUBAH: Membuat konten berkumpul di tengah */
  max-width: 1140px; /* Menyelaraskan lebar dengan konten atas website */
  margin: 0 auto;
  padding: 0 40px; /* Jarak aman agar teks tidak terlalu menempel ke pinggir layar */
}

/* Kolom Kiri Footer */
.footer-left-content {
  display: none; /* DIUBAH: Disembunyikan karena kosong, agar tidak memakan space */
}

/* Kolom Kanan Footer (Teks Copyright) */
.footer-copyright {
  text-align: center; /* DIUBAH: Memastikan teks rata tengah */
  width: 100%; /* Memastikan block mengambil ruang penuh untuk fungsi centering */
}

.footer-copyright p {
  font-family: Arial, sans-serif;
  font-size: 13px; /* Ukuran teks kecil, tipis, dan normal untuk footer */
  color: #ffffff;  /* Warna teks putih bersih agar kontras dengan background hijau */
  margin: 0;
  font-weight: normal;
  letter-spacing: 0.3px;
}
</style>
