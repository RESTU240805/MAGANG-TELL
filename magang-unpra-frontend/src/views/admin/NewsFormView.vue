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

      <div class="grid grid-cols-3 gap-8">
        <div class="col-span-2 space-y-6">

          <!-- Form Info -->
          <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100">
            <h3 class="font-bold text-gray-900 mb-6">Informasi Berita</h3>
            <div class="space-y-5">
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
                <label class="text-sm font-semibold text-gray-700 block mb-2">Summary</label>
                <input v-model="form.summary" type="text" placeholder="Ringkasan singkat berita..."
                  class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
              </div>
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-2">Konten</label>
                <textarea v-model="form.content" rows="8" placeholder="Tulis isi berita di sini..."
                  class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 resize-none"></textarea>
              </div>
              <div class="flex items-center gap-3">
                <input type="checkbox" v-model="form.is_published" id="published" class="w-4 h-4 accent-green-600" />
                <label for="published" class="text-sm font-semibold text-gray-700">Publish berita ini</label>
              </div>
            </div>
          </div>

          <!-- Upload Gambar -->
          <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100">
            <h3 class="font-bold text-gray-900 mb-2">Gambar Berita</h3>
            <p class="text-gray-400 text-xs mb-6">Upload gambar untuk berita ini. Bisa lebih dari satu.</p>

            <!-- Existing Images -->
            <div v-if="imageList.length > 0" class="grid grid-cols-3 gap-4 mb-6">
              <div v-for="(img, i) in imageList" :key="i"
                class="relative group rounded-xl overflow-hidden h-36 bg-gray-100 border border-gray-200">
                <img :src="img.url" :alt="`Image ${i+1}`" class="w-full h-full object-cover" />
                <div class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition flex items-center justify-center">
                  <button @click="removeImage(i)"
                    class="bg-red-500 text-white px-3 py-1.5 rounded-lg text-xs font-semibold">
                    🗑 Hapus
                  </button>
                </div>
              </div>
            </div>

            <!-- Upload Area -->
            <div class="border-2 border-dashed border-gray-200 rounded-2xl p-8 text-center hover:border-green-400 transition cursor-pointer"
              @click="triggerFileInput" @dragover.prevent @drop.prevent="handleDrop">
              <input ref="fileInput" type="file" accept="image/*" multiple class="hidden" @change="handleFileChange" />
              <div v-if="uploading" class="flex flex-col items-center gap-3">
                <div class="w-8 h-8 border-4 border-green-600 border-t-transparent rounded-full animate-spin"></div>
                <p class="text-sm text-gray-500">Mengupload gambar...</p>
              </div>
              <div v-else class="flex flex-col items-center gap-3">
                <div class="w-14 h-14 bg-green-50 rounded-2xl flex items-center justify-center text-3xl">🖼️</div>
                <div>
                  <p class="font-semibold text-gray-700 text-sm">Klik untuk upload atau drag & drop</p>
                  <p class="text-gray-400 text-xs mt-1">PNG, JPG, WebP — Maksimal 5MB</p>
                </div>
                <button type="button"
                  class="bg-green-600 text-white px-6 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
                  Pilih Gambar
                </button>
              </div>
            </div>
          </div>

          <p v-if="error" class="text-red-500 text-sm bg-red-50 px-4 py-3 rounded-xl">{{ error }}</p>

          <div class="flex gap-4">
            <button @click="handleSubmit" :disabled="loading || uploading"
              class="bg-green-600 text-white px-8 py-3 rounded-xl font-semibold hover:bg-green-700 transition disabled:opacity-50">
              {{ loading ? 'Menyimpan...' : (isEdit ? 'Update Berita' : 'Simpan Berita') }}
            </button>
            <RouterLink to="/admin/news"
              class="border border-gray-200 px-8 py-3 rounded-xl font-semibold text-gray-600 hover:bg-gray-50 transition">
              Batal
            </RouterLink>
          </div>
        </div>

        <!-- Sidebar -->
        <div class="space-y-6">
          <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
            <h3 class="font-bold text-gray-900 mb-4">📋 Panduan</h3>
            <ul class="text-sm text-gray-500 space-y-3">
              <li class="flex items-start gap-2"><span class="text-green-500">✓</span> Isi judul dan konten</li>
              <li class="flex items-start gap-2"><span class="text-green-500">✓</span> Upload gambar berita</li>
              <li class="flex items-start gap-2"><span class="text-green-500">✓</span> Centang Publish agar tampil di website</li>
              <li class="flex items-start gap-2"><span class="text-yellow-500">⚠</span> Kalau tidak dicentang, berita tidak tampil di publik</li>
            </ul>
          </div>
          <div class="bg-green-50 rounded-2xl p-6 border border-green-100 text-center">
            <p class="text-4xl font-black text-green-600">{{ imageList.length }}</p>
            <p class="text-sm text-green-700 mt-1">Gambar diupload</p>
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
const uploading = ref(false)
const error = ref('')
const fileInput = ref(null)
const imageList = ref([])

const form = ref({
  title: '',
  category: '',
  summary: '',
  content: '',
  is_published: false,
})

onMounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }

  if (isEdit.value) {
    const res = await api.get(`/news/${route.params.id}`)
    const data = res.data.data
    form.value = {
      title: data.title,
      category: data.category,
      summary: data.summary || '',
      content: data.content,
      is_published: data.is_published,
    }
    if (data.Images) {
      imageList.value = data.Images.map(img => ({ url: img.image_url, id: img.ID }))
    }
  }
})

const triggerFileInput = () => fileInput.value.click()

const handleFileChange = async (e) => {
  await uploadFiles(Array.from(e.target.files))
}

const handleDrop = async (e) => {
  const files = Array.from(e.dataTransfer.files).filter(f => f.type.startsWith('image/'))
  await uploadFiles(files)
}

const uploadFiles = async (files) => {
  if (!files.length) {return}
  uploading.value = true
  error.value = ''
  for (const file of files) {
    if (file.size > 5 * 1024 * 1024) {
      error.value = `File ${file.name} terlalu besar (maks 5MB)`
      continue
    }
    try {
      const formData = new FormData()
      formData.append('image', file)
      const res = await api.post('/upload', formData, {
        headers: { 'Content-Type': 'multipart/form-data' }
      })
      imageList.value.push({ url: res.data.url })
    } catch {
      error.value = `Gagal upload ${file.name}`
    }
  }
  uploading.value = false
  if (fileInput.value) {fileInput.value.value = ''}
}

const removeImage = (index) => imageList.value.splice(index, 1)

const handleSubmit = async () => {
  if (!form.value.title || !form.value.content) {
    error.value = 'Judul dan konten wajib diisi!'
    return
  }
  loading.value = true
  error.value = ''
  try {
    const payload = {
      ...form.value,
      Images: imageList.value.map(img => ({ image_url: img.url }))
    }
    if (isEdit.value) {
      await api.put(`/admin/news/${route.params.id}`, payload)
    } else {
      await api.post('/admin/news', payload)
    }
    router.push('/admin/news')
  } catch (_err) {
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