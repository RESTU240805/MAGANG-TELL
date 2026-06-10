<template>
  <div class="min-h-screen bg-gray-100 flex">

    <!-- Sidebar (sama) -->
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
      <nav class="flex-1 p-4 space-y-1">
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-2 px-2">OVERVIEW</p>
        <RouterLink to="/admin/dashboard"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📊 Dashboard
        </RouterLink>
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-4 px-2">CONTENT ENGINE</p>
        <RouterLink to="/admin/news"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
          📰 Corporate News
        </RouterLink>
        <RouterLink to="/admin/products"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📦 Products
        </RouterLink>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <button @click="logout" class="text-sm text-gray-400 hover:text-white transition px-2">→ Logout</button>
      </div>
    </aside>

    <!-- Main -->
    <main class="flex-1 ml-64 p-10">
      <div class="flex justify-between items-center mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Corporate News</h1>
        </div>
        <RouterLink to="/admin/news/create"
          class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
          + Tambah Berita
        </RouterLink>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="text-center text-gray-400 py-20">Loading...</div>

      <!-- Empty -->
      <div v-else-if="newsList.length === 0"
        class="bg-white rounded-2xl p-16 text-center border border-gray-100">
        <p class="text-4xl mb-4">📰</p>
        <p class="text-gray-500">Belum ada berita. Tambahkan berita pertama!</p>
        <RouterLink to="/admin/news/create"
          class="mt-4 inline-block bg-green-600 text-white px-6 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
          + Tambah Berita
        </RouterLink>
      </div>

      <!-- News Table -->
      <div v-else class="bg-white rounded-2xl shadow-sm border border-gray-100 overflow-hidden">
        <table class="w-full">
          <thead class="bg-gray-50 border-b border-gray-100">
            <tr>
              <th class="text-left px-6 py-4 text-xs font-semibold text-gray-500 tracking-widest">JUDUL</th>
              <th class="text-left px-6 py-4 text-xs font-semibold text-gray-500 tracking-widest">KATEGORI</th>
              <th class="text-left px-6 py-4 text-xs font-semibold text-gray-500 tracking-widest">TANGGAL</th>
              <th class="text-left px-6 py-4 text-xs font-semibold text-gray-500 tracking-widest">AKSI</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-50">
            <tr v-for="news in newsList" :key="news.ID" class="hover:bg-gray-50 transition">
              <td class="px-6 py-4">
                <p class="font-semibold text-gray-900 text-sm">{{ news.title }}</p>
                <p class="text-gray-400 text-xs mt-1 line-clamp-1">{{ news.content }}</p>
              </td>
              <td class="px-6 py-4">
                <span class="bg-green-100 text-green-700 px-3 py-1 rounded-full text-xs font-semibold">
                  {{ news.category || 'Umum' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ new Date(news.CreatedAt).toLocaleDateString('id-ID') }}
              </td>
              <td class="px-6 py-4">
                <div class="flex gap-2">
                  <RouterLink :to="`/admin/news/edit/${news.ID}`"
                    class="bg-blue-50 text-blue-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-blue-100 transition">
                    Edit
                  </RouterLink>
                  <button @click="deleteNews(news.ID)"
                    class="bg-red-50 text-red-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-red-100 transition">
                    Hapus
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const router = useRouter()
const newsList = ref([])
const loading = ref(true)

onMounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }
  try {
    const res = await api.get('/news')
    newsList.value = res.data.data
  } finally {
    loading.value = false
  }
})

const deleteNews = async (id) => {
  if (!confirm('Yakin hapus berita ini?')) return
  await api.delete(`/news/${id}`)
  newsList.value = newsList.value.filter(n => n.ID !== id)
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}
</script>