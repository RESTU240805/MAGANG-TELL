<template>
  <div class="min-h-screen bg-gray-100 flex">

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

    <main class="flex-1 ml-64 p-10">
      <div class="flex items-center gap-4 mb-8">
        <RouterLink to="/admin/news" class="text-gray-400 hover:text-gray-600 transition">← Kembali</RouterLink>
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">{{ isEdit ? 'Edit Berita' : 'Tambah Berita' }}</h1>
        </div>
      </div>

      <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100 max-w-3xl">
        <div class="space-y-6">
          <div>
            <label class="text-sm font-semibold text-gray-700 block mb-2">Judul Berita</label>
            <input v-model="form.title" type="text" placeholder="Masukkan judul berita..."
              class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
          </div>
          <div>
            <label class="text-sm font-semibold text-gray-700 block mb-2">Kategori</label>
            <input v-model="form.category" type="text" placeholder="Contoh: Operational Excellence"
              class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
          </div>
          <div>
            <label class="text-sm font-semibold text-gray-700 block mb-2">Konten</label>
            <textarea v-model="form.content" rows="8" placeholder="Tulis isi berita di sini..."
              class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 resize-none"></textarea>
          </div>

          <p v-if="error" class="text-red-500 text-sm">{{ error }}</p>

          <div class="flex gap-4">
            <button @click="handleSubmit"
              class="bg-green-600 text-white px-8 py-3 rounded-xl font-semibold hover:bg-green-700 transition">
              {{ loading ? 'Menyimpan...' : (isEdit ? 'Update Berita' : 'Simpan Berita') }}
            </button>
            <RouterLink to="/admin/news"
              class="border border-gray-200 px-8 py-3 rounded-xl font-semibold text-gray-600 hover:bg-gray-50 transition">
              Batal
            </RouterLink>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute, RouterLink } from 'vue-router'
import api from '../../services/api'

const router = useRouter()
const route = useRoute()
const isEdit = computed(() => !!route.params.id)
const loading = ref(false)
const error = ref('')

const form = ref({ title: '', category: '', content: '' })

onMounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }

  if (isEdit.value) {
    const res = await api.get(`/news/${route.params.id}`)
    form.value = {
      title: res.data.data.title,
      category: res.data.data.category,
      content: res.data.data.content,
    }
  }
})

const handleSubmit = async () => {
  if (!form.value.title || !form.value.content) {
    error.value = 'Judul dan konten wajib diisi!'
    return
  }
  loading.value = true
  try {
    if (isEdit.value) {
      await api.put(`/news/${route.params.id}`, form.value)
    } else {
      await api.post('/news', form.value)
    }
    router.push('/admin/news')
  } catch (err) {
    error.value = 'Gagal menyimpan berita.'
  } finally {
    loading.value = false
  }
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}
</script>