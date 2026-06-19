<template>
  <div class="min-h-screen bg-gray-100 flex">
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
      <nav class="flex-1 p-4 space-y-1">
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-2 px-2">OVERVIEW</p>
        <RouterLink to="/admin/dashboard"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📊 Dashboard
        </RouterLink>
        <p class="text-xs text-gray-500 font-semibold tracking-widest mb-2 mt-4 px-2">CONTENT ENGINE</p>
        <RouterLink to="/admin/news"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📰 Corporate News
        </RouterLink>
        <RouterLink to="/admin/products"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📦 Products
        </RouterLink>
        <RouterLink to="/admin/slider"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🖼️ Product Slider
        </RouterLink>
        <RouterLink to="/admin/product-page"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
          📝 Product Page
        </RouterLink>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <button @click="logout" class="text-sm text-gray-400 hover:text-white transition px-2">→ Logout</button>
      </div>
    </aside>

    <main class="flex-1 ml-64 p-10">
      <div class="mb-8">
        <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
        <h1 class="text-3xl font-black text-gray-900">Product Page</h1>
        <p class="text-gray-400 text-sm mt-1">Kelola deskripsi utama yang tampil di halaman produk.</p>
      </div>

      <div v-if="alert.show"
        :class="['mb-4 px-4 py-3 rounded-xl text-sm font-medium flex items-center gap-2',
          alert.type === 'success' ? 'bg-green-50 text-green-700 border border-green-200' : 'bg-red-50 text-red-700 border border-red-200']">
        {{ alert.type === 'success' ? '✅' : '❌' }} {{ alert.message }}
      </div>

      <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100 max-w-3xl">
        <div v-if="loading" class="flex items-center justify-center py-16 text-gray-400">
          <div class="w-5 h-5 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
        </div>
        <div v-else class="space-y-5">
          <div>
            <label class="text-sm font-semibold text-gray-700 block mb-2">Deskripsi Produk</label>
            <textarea v-model="description" rows="16" placeholder="Tulis deskripsi detail produk pulp di sini..."
              class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 resize-none"></textarea>
            <p class="text-xs text-gray-400 mt-2">Deskripsi ini akan tampil di halaman /product, menggantikan katalog produk.</p>
          </div>

          <div class="flex gap-4 pt-2">
            <button @click="save" :disabled="saving"
              class="bg-green-600 text-white px-8 py-3 rounded-xl font-semibold hover:bg-green-700 transition disabled:opacity-50 flex items-center gap-2">
              <div v-if="saving" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
              {{ saving ? 'Menyimpan...' : 'Simpan' }}
            </button>
          </div>
        </div>
      </div>

      <!-- Preview -->
      <div v-if="description" class="mt-8 max-w-3xl">
        <h3 class="text-sm font-semibold text-gray-500 tracking-widest mb-3">PRATINJAU</h3>
        <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100 prose prose-sm max-w-none">
          <div class="whitespace-pre-wrap text-gray-700 leading-relaxed">{{ description }}</div>
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
const description = ref('')
const loading = ref(true)
const saving = ref(false)
const alert = ref({ show: false, type: 'success', message: '' })

const showAlert = (type, message) => {
  alert.value = { show: true, type, message }
  setTimeout(() => { alert.value.show = false }, 3000)
}

const fetch = async () => {
  loading.value = true
  try {
    const res = await api.get('/product-page')
    description.value = res.data?.data?.description || ''
  } catch { showAlert('error', 'Gagal memuat data') }
  finally { loading.value = false }
}

const save = async () => {
  saving.value = true
  try {
    await api.put('/product-page', { description: description.value })
    showAlert('success', 'Deskripsi produk berhasil disimpan!')
  } catch (e) {
    console.error('Save error:', e)
    showAlert('error', e.response?.data?.error || e.message || 'Gagal menyimpan')
  }
  finally { saving.value = false }
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}

onMounted(() => {
  if (!localStorage.getItem('token')) { router.push('/admin/login'); return }
  fetch()
})
</script>
