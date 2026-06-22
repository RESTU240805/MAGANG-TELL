<template>
  <div class="min-h-screen bg-gray-100 flex">

    <!-- Sidebar -->
    <aside class="w-64 bg-gray-950 text-white flex flex-col fixed h-full z-40">
      <div class="p-6 border-b border-gray-800">
        <div class="flex items-center gap-3">
          <div class="w-9 h-9 bg-green-600 rounded-xl flex items-center justify-center font-black text-sm">T</div>
          <div>
            <p class="font-black text-sm">TELPP</p>
            <p class="text-xs text-gray-500">Management Profile</p>
          </div>
        </div>
      </div>

      <nav class="flex-1 p-4 space-y-1 overflow-y-auto">
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-2 px-2">OVERVIEW</p>
        <RouterLink to="/admin/dashboard"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
          📊 Dashboard
        </RouterLink>
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-4 px-2">CONTENT ENGINE</p>
        <RouterLink to="/admin/news"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📰 Corporate News
        </RouterLink>
        <RouterLink to="/admin/product-page"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📦 Product
        </RouterLink>
        <RouterLink to="/admin/about"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏢 About Section
        </RouterLink>
        <RouterLink to="/admin/community"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🌱 Community
        </RouterLink>
        <RouterLink to="/admin/team-members"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          👥 Our Team
        </RouterLink>
        <RouterLink to="/admin/our-company"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏛️ Our Company
        </RouterLink>
        <RouterLink to="/admin/menus"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📋 Menu
        </RouterLink>
      </nav>

      <div class="p-4 border-t border-gray-800">
        <div class="flex items-center gap-3 mb-3">
          <div class="w-8 h-8 bg-green-600 rounded-full flex items-center justify-center text-xs font-bold">A</div>
          <div>
            <p class="text-sm font-medium">{{ user.name }}</p>
            <p class="text-xs text-gray-500">{{ user.email }}</p>
          </div>
        </div>
        <button @click="logout"
          class="w-full text-left text-sm text-gray-400 hover:text-white transition px-2">
          → Logout
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <main class="flex-1 ml-64 p-8">

      <!-- Header -->
      <div class="flex justify-between items-center mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">ADMIN SYSTEM</p>
          <h1 class="text-3xl font-black text-gray-900">Dashboard</h1>
        </div>
        <div class="text-right">
          <p class="text-sm font-semibold text-gray-700">{{ currentDay }}</p>
          <p class="text-xs text-gray-400">{{ currentDate }}</p>
        </div>
      </div>

      <!-- Stats Grid -->
      <div class="grid grid-cols-2 lg:grid-cols-4 gap-5 mb-8">
        <div class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100 hover:shadow-md transition">
          <div class="flex items-center justify-between mb-3">
            <div class="w-10 h-10 bg-blue-50 rounded-xl flex items-center justify-center text-lg">📰</div>
            <span class="text-xs text-green-600 font-semibold bg-green-50 px-2 py-0.5 rounded-full">Aktif</span>
          </div>
          <p class="text-3xl font-black text-gray-900">{{ newsCount }}</p>
          <p class="text-xs text-gray-400 mt-1">Berita Dipublikasi</p>
        </div>
        <div class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100 hover:shadow-md transition">
          <div class="flex items-center justify-between mb-3">
            <div class="w-10 h-10 bg-purple-50 rounded-xl flex items-center justify-center text-lg">📦</div>
            <span class="text-xs text-green-600 font-semibold bg-green-50 px-2 py-0.5 rounded-full">Aktif</span>
          </div>
          <p class="text-3xl font-black text-gray-900">{{ productCount }}</p>
          <p class="text-xs text-gray-400 mt-1">Produk Terdaftar</p>
        </div>
        <div class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100 hover:shadow-md transition">
          <div class="flex items-center justify-between mb-3">
            <div class="w-10 h-10 bg-amber-50 rounded-xl flex items-center justify-center text-lg">🖼️</div>
            <span class="text-xs text-blue-600 font-semibold bg-blue-50 px-2 py-0.5 rounded-full">Slide</span>
          </div>
          <p class="text-3xl font-black text-gray-900">{{ slideCount }}</p>
          <p class="text-xs text-gray-400 mt-1">Slider Aktif</p>
        </div>
        <div class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100 hover:shadow-md transition">
          <div class="flex items-center justify-between mb-3">
            <div class="w-10 h-10 bg-green-50 rounded-xl flex items-center justify-center text-lg">👥</div>
            <span class="text-xs text-blue-600 font-semibold bg-blue-50 px-2 py-0.5 rounded-full">Tim</span>
          </div>
          <p class="text-3xl font-black text-gray-900">{{ teamCount }}</p>
          <p class="text-xs text-gray-400 mt-1">Anggota Tim</p>
        </div>
      </div>

      <!-- Welcome Banner -->
      <div class="bg-gradient-to-r from-gray-900 to-gray-800 rounded-2xl p-8 text-white mb-8 relative overflow-hidden">
        <div class="absolute -right-10 -top-10 w-40 h-40 bg-green-600/10 rounded-full blur-2xl"></div>
        <div class="absolute -right-5 -bottom-10 w-32 h-32 bg-green-600/10 rounded-full blur-2xl"></div>
        <div class="relative flex items-center gap-5">
          <div class="w-14 h-14 bg-green-600/20 rounded-2xl flex items-center justify-center text-3xl shrink-0">👋</div>
          <div>
            <h2 class="text-xl font-black mb-1">Selamat Datang, {{ user.name }}</h2>
            <p class="text-gray-400 text-sm leading-relaxed">Kelola konten website perusahaan dari sini. Anda bisa mengupdate berita, produk, slider, profil perusahaan, dan menu navigasi.</p>
          </div>
        </div>
      </div>

      <!-- Bottom Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-5">

        <!-- Quick Actions -->
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <h3 class="font-bold text-gray-900 mb-4 flex items-center gap-2">
            <span class="w-2 h-2 bg-green-500 rounded-full"></span>
            Quick Actions
          </h3>
          <div class="space-y-2">
            <RouterLink to="/admin/news/create"
              class="flex items-center gap-3 p-3 rounded-xl hover:bg-gray-50 transition group">
              <div class="w-9 h-9 bg-blue-50 rounded-lg flex items-center justify-center text-sm group-hover:bg-blue-100 transition">✏️</div>
              <div>
                <p class="text-sm font-semibold text-gray-700">Tambah Berita</p>
                <p class="text-xs text-gray-400">Publikasikan pengumuman baru</p>
              </div>
            </RouterLink>
            <RouterLink to="/admin/products/create"
              class="flex items-center gap-3 p-3 rounded-xl hover:bg-gray-50 transition group">
              <div class="w-9 h-9 bg-purple-50 rounded-lg flex items-center justify-center text-sm group-hover:bg-purple-100 transition">📦</div>
              <div>
                <p class="text-sm font-semibold text-gray-700">Tambah Produk</p>
                <p class="text-xs text-gray-400">Daftarkan produk baru</p>
              </div>
            </RouterLink>
            <RouterLink to="/admin/slider"
              class="flex items-center gap-3 p-3 rounded-xl hover:bg-gray-50 transition group">
              <div class="w-9 h-9 bg-amber-50 rounded-lg flex items-center justify-center text-sm group-hover:bg-amber-100 transition">🖼️</div>
              <div>
                <p class="text-sm font-semibold text-gray-700">Kelola Slider</p>
                <p class="text-xs text-gray-400">Atur gambar hero produk</p>
              </div>
            </RouterLink>
            <RouterLink to="/admin/our-company"
              class="flex items-center gap-3 p-3 rounded-xl hover:bg-gray-50 transition group">
              <div class="w-9 h-9 bg-green-50 rounded-lg flex items-center justify-center text-sm group-hover:bg-green-100 transition">🏛️</div>
              <div>
                <p class="text-sm font-semibold text-gray-700">Profil Perusahaan</p>
                <p class="text-xs text-gray-400">Update hero, creed & dokumen</p>
              </div>
            </RouterLink>
            <RouterLink to="/admin/menus"
              class="flex items-center gap-3 p-3 rounded-xl hover:bg-gray-50 transition group">
              <div class="w-9 h-9 bg-gray-100 rounded-lg flex items-center justify-center text-sm group-hover:bg-gray-200 transition">📋</div>
              <div>
                <p class="text-sm font-semibold text-gray-700">Kelola Menu</p>
                <p class="text-xs text-gray-400">Aktifkan/nonaktifkan menu navbar</p>
              </div>
            </RouterLink>
          </div>
        </div>

        <!-- Content Overview -->
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <h3 class="font-bold text-gray-900 mb-4 flex items-center gap-2">
            <span class="w-2 h-2 bg-blue-500 rounded-full"></span>
            Ringkasan Konten
          </h3>
          <div class="space-y-3">
            <RouterLink to="/admin/news" class="flex items-center justify-between p-3 rounded-xl hover:bg-gray-50 transition">
              <div class="flex items-center gap-3">
                <div class="w-9 h-9 bg-blue-50 rounded-lg flex items-center justify-center text-sm">📰</div>
                <span class="text-sm font-medium text-gray-700">Berita</span>
              </div>
              <span class="text-sm font-bold text-gray-900">{{ newsCount }}</span>
            </RouterLink>
            <RouterLink to="/admin/products" class="flex items-center justify-between p-3 rounded-xl hover:bg-gray-50 transition">
              <div class="flex items-center gap-3">
                <div class="w-9 h-9 bg-purple-50 rounded-lg flex items-center justify-center text-sm">📦</div>
                <span class="text-sm font-medium text-gray-700">Produk</span>
              </div>
              <span class="text-sm font-bold text-gray-900">{{ productCount }}</span>
            </RouterLink>
            <RouterLink to="/admin/slider" class="flex items-center justify-between p-3 rounded-xl hover:bg-gray-50 transition">
              <div class="flex items-center gap-3">
                <div class="w-9 h-9 bg-amber-50 rounded-lg flex items-center justify-center text-sm">🖼️</div>
                <span class="text-sm font-medium text-gray-700">Slider</span>
              </div>
              <span class="text-sm font-bold text-gray-900">{{ slideCount }}</span>
            </RouterLink>
            <RouterLink to="/admin/team-members" class="flex items-center justify-between p-3 rounded-xl hover:bg-gray-50 transition">
              <div class="flex items-center gap-3">
                <div class="w-9 h-9 bg-green-50 rounded-lg flex items-center justify-center text-sm">👥</div>
                <span class="text-sm font-medium text-gray-700">Tim</span>
              </div>
              <span class="text-sm font-bold text-gray-900">{{ teamCount }}</span>
            </RouterLink>
            <RouterLink to="/admin/about" class="flex items-center justify-between p-3 rounded-xl hover:bg-gray-50 transition">
              <div class="flex items-center gap-3">
                <div class="w-9 h-9 bg-pink-50 rounded-lg flex items-center justify-center text-sm">🏢</div>
                <span class="text-sm font-medium text-gray-700">About</span>
              </div>
              <span class="text-xs font-medium text-gray-500 bg-gray-100 px-2 py-0.5 rounded-full">Section</span>
            </RouterLink>
          </div>
        </div>

        <!-- System Status -->
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <h3 class="font-bold text-gray-900 mb-4 flex items-center gap-2">
            <span class="w-2 h-2 bg-green-500 rounded-full"></span>
            Status Sistem
          </h3>
          <div class="space-y-3">
            <div class="flex items-center justify-between p-3 rounded-xl bg-green-50 border border-green-100">
              <div class="flex items-center gap-3">
                <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
                <span class="text-sm font-medium text-gray-700">Backend Golang</span>
              </div>
              <span class="text-xs font-semibold text-green-600 bg-green-100 px-2 py-0.5 rounded-full">Online</span>
            </div>
            <div class="flex items-center justify-between p-3 rounded-xl bg-green-50 border border-green-100">
              <div class="flex items-center gap-3">
                <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
                <span class="text-sm font-medium text-gray-700">PostgreSQL</span>
              </div>
              <span class="text-xs font-semibold text-green-600 bg-green-100 px-2 py-0.5 rounded-full">Connected</span>
            </div>
            <div class="flex items-center justify-between p-3 rounded-xl bg-green-50 border border-green-100">
              <div class="flex items-center gap-3">
                <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
                <span class="text-sm font-medium text-gray-700">Frontend Vue</span>
              </div>
              <span class="text-xs font-semibold text-green-600 bg-green-100 px-2 py-0.5 rounded-full">Running</span>
            </div>
          </div>
          <div class="mt-5 p-4 bg-gray-50 rounded-xl">
            <p class="text-xs text-gray-400 mb-1">Versi Aplikasi</p>
            <p class="text-sm font-semibold text-gray-700">TELPP Management Profile v1.0</p>
          </div>
        </div>
      </div>

    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const router = useRouter()
const user = ref({ name: 'Administrator', email: '' })
const newsCount = ref(0)
const productCount = ref(0)
const slideCount = ref(0)
const teamCount = ref(0)

const days = ['Minggu', 'Senin', 'Selasa', 'Rabu', 'Kamis', 'Jumat', 'Sabtu']
const now = new Date()
const currentDay = days[now.getDay()]
const currentDate = now.toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' })

onMounted(async () => {
  const stored = localStorage.getItem('user')
  if (stored) { user.value = JSON.parse(stored) }

  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }

  try {
    const [newsRes, productRes, slideRes, teamRes] = await Promise.allSettled([
      api.get('/admin/news'),
      api.get('/admin/products'),
      api.get('/admin/product-slides'),
      api.get('/admin/team-members'),
    ])
    if (newsRes.status === 'fulfilled') { newsCount.value = (newsRes.value.data?.data || newsRes.value.data || []).length }
    if (productRes.status === 'fulfilled') { productCount.value = (productRes.value.data?.data || productRes.value.data || []).length }
    if (slideRes.status === 'fulfilled') { slideCount.value = (slideRes.value.data?.data || slideRes.value.data || []).length }
    if (teamRes.status === 'fulfilled') { teamCount.value = (teamRes.value.data?.data || teamRes.value.data || []).length }
  } catch { /* silent */ }
})

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}
</script>
