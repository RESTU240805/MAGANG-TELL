<template>
  <div class="min-h-screen bg-gray-100 flex">

    <!-- Sidebar -->
    <aside class="w-64 bg-gray-950 text-white flex flex-col fixed h-full">
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
        <RouterLink to="/admin/slider"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🖼️ Product Slider
        </RouterLink>
        <RouterLink to="/admin/product-page"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📝 Product Page
        </RouterLink>
        <RouterLink to="/admin/about"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏢 About Section
        </RouterLink>
        <RouterLink to="/admin/team-members"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          👥 Our Team
        </RouterLink>
        <RouterLink to="/admin/our-company"
        class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
         🏛️ Our Company
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
    <main class="flex-1 ml-64 p-10">

      <!-- Header -->
      <div class="flex justify-between items-start mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">ADMIN SYSTEM</p>
          <h1 class="text-3xl font-black text-gray-900">Dashboard Overview</h1>
        </div>
        <div class="text-right">
          <p class="text-sm font-semibold text-gray-700">{{ currentDay }}</p>
          <p class="text-xs text-gray-400">{{ currentDate }}</p>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-4 gap-6 mb-8">
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <p class="text-xs text-gray-400 tracking-widest font-semibold">PUBLISHED NEWS</p>
          <p class="text-4xl font-black text-gray-900 mt-2">{{ newsCount }}</p>
          <p class="text-gray-400 text-sm mt-1">Current active announcements.</p>
        </div>
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <p class="text-xs text-gray-400 tracking-widest font-semibold">PRODUCTS</p>
          <p class="text-4xl font-black text-gray-900 mt-2">{{ productCount }}</p>
          <p class="text-gray-400 text-sm mt-1">Pulp mill product categories.</p>
        </div>
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <p class="text-xs text-gray-400 tracking-widest font-semibold">TOTAL VISITS</p>
          <p class="text-4xl font-black text-gray-900 mt-2">12.8K</p>
          <p class="text-green-500 text-sm mt-1">↑ +14.5% last mo.</p>
        </div>
        <div class="bg-green-600 rounded-2xl p-6 shadow-sm text-white">
          <p class="text-xs tracking-widest font-semibold opacity-70">SYSTEM STATUS</p>
          <p class="text-3xl font-black mt-2">Secure</p>
          <p class="text-sm mt-1 opacity-70">Everything is running smoothly.</p>
        </div>
      </div>

      <!-- Welcome Banner -->
      <div class="bg-gray-900 rounded-2xl p-8 text-white mb-8">
        <div class="flex items-center gap-6">
          <div class="w-16 h-16 bg-green-600/20 rounded-full flex items-center justify-center text-3xl">👋</div>
          <div>
            <h2 class="text-2xl font-black mb-1">Welcome back, {{ user.name }}.</h2>
            <p class="text-gray-400">Use the sidebar to manage your company's digital narrative. You can update historical records, modify the company creed, or publish new updates.</p>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="grid grid-cols-2 gap-6">
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <h3 class="font-bold text-gray-900 mb-4">Quick Actions</h3>
          <div class="space-y-3">
            <RouterLink to="/admin/news/create"
              class="flex items-center gap-3 p-3 rounded-xl bg-green-50 hover:bg-green-100 transition text-green-700 text-sm font-medium">
              ✏️ Tambah Berita Baru
            </RouterLink>
            <RouterLink to="/admin/products/create"
              class="flex items-center gap-3 p-3 rounded-xl bg-blue-50 hover:bg-blue-100 transition text-blue-700 text-sm font-medium">
              📦 Tambah Produk Baru
            </RouterLink>
          </div>
        </div>
        <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
          <h3 class="font-bold text-gray-900 mb-4">Info Sistem</h3>
          <div class="space-y-2 text-sm text-gray-600">
            <p>🟢 Backend Golang: <span class="font-semibold text-green-600">Online</span></p>
            <p>🟢 Database PostgreSQL: <span class="font-semibold text-green-600">Connected</span></p>
            <p>🟢 Frontend Vue: <span class="font-semibold text-green-600">Running</span></p>
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

const days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']
const now = new Date()
const currentDay = days[now.getDay()]
const currentDate = now.toLocaleDateString('en-US', { month: 'long', day: 'numeric', year: 'numeric' })

onMounted(async () => {
  const stored = localStorage.getItem('user')
  if (stored) {user.value = JSON.parse(stored)}

  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }

  const [newsRes, productRes] = await Promise.all([
    api.get('/admin/news'),
    api.get('/admin/products')
  ])
  newsCount.value = newsRes.data.data.length
  productCount.value = productRes.data.data.length
})

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}
</script>