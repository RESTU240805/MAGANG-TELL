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
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
          📦 Products
        </RouterLink>
      </nav>
      <div class="p-4 border-t border-gray-800">
        <button @click="logout" class="text-sm text-gray-400 hover:text-white transition px-2">→ Logout</button>
      </div>
    </aside>

    <!-- Main -->
    <main class="flex-1 ml-64 p-10">
      <div class="flex items-center gap-4 mb-8">
        <RouterLink to="/admin/products" class="text-gray-400 hover:text-gray-600 transition">← Kembali</RouterLink>
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">{{ isEdit ? 'Edit Produk' : 'Tambah Produk' }}</h1>
        </div>
      </div>

      <div class="grid grid-cols-3 gap-8">

        <!-- Form -->
        <div class="col-span-2 space-y-6">
          <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100">
            <h3 class="font-bold text-gray-900 mb-6">Informasi Produk</h3>
            <div class="space-y-5">
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-2">Nama Produk</label>
                <input v-model="form.name" type="text" placeholder="Contoh: TeL Pellita Bleached Kraft Pulp"
                  class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
              </div>
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-2">Deskripsi</label>
                <textarea v-model="form.description" rows="6"
                  placeholder="Masukkan deskripsi produk..."
                  class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 resize-none"></textarea>
              </div>
            </div>
          </div>

          <!-- Image Upload -->
          <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100">
            <h3 class="font-bold text-gray-900 mb-2">Gambar Produk</h3>
            <p class="text-gray-400 text-xs mb-6">Upload gambar dari komputer Anda. Bisa lebih dari satu gambar.</p>

            <!-- Existing Images -->
            <div v-if="imageList.length > 0" class="grid grid-cols-3 gap-4 mb-6">
              <div v-for="(img, i) in imageList" :key="i"
                class="relative group rounded-xl overflow-hidden h-32 bg-gray-100">
                <img :src="getImageSrc(img.url)" :alt="`Image ${i+1}`" class="w-full h-full object-cover" />
                <!-- Upload progress overlay -->
                <div v-if="img.uploading"
                  class="absolute inset-0 bg-black/60 flex flex-col items-center justify-center gap-2">
                  <div class="w-8 h-8 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                  <span class="text-white text-xs">Mengupload...</span>
                </div>
                <!-- Hover delete -->
                <div v-else class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition flex items-center justify-center">
                  <button @click="removeImage(i)"
                    class="bg-red-500 text-white px-3 py-1 rounded-lg text-xs font-semibold">
                    Hapus
                  </button>
                </div>
              </div>
            </div>

            <!-- Drop Zone / File Input -->
            <div
              class="border-2 border-dashed border-gray-200 rounded-xl p-8 text-center hover:border-green-400 hover:bg-green-50 transition cursor-pointer"
              :class="{ 'border-green-400 bg-green-50': isDragging }"
              @click="triggerFileInput"
              @dragover.prevent="isDragging = true"
              @dragleave.prevent="isDragging = false"
              @drop.prevent="handleDrop">
              <div class="text-4xl mb-3">🖼️</div>
              <p class="text-sm font-semibold text-gray-700">Klik atau seret gambar ke sini</p>
              <p class="text-xs text-gray-400 mt-1">JPG, PNG, WebP — maks. 5MB per file</p>
              <input
                ref="fileInput"
                type="file"
                accept="image/jpeg,image/png,image/webp"
                multiple
                class="hidden"
                @change="handleFileChange" />
            </div>

            <p v-if="uploadError" class="text-red-500 text-sm mt-3">{{ uploadError }}</p>
          </div>

          <p v-if="error" class="text-red-500 text-sm">{{ error }}</p>

          <div class="flex gap-4">
            <button @click="handleSubmit"
              :disabled="loading || imageList.some(i => i.uploading)"
              class="bg-green-600 text-white px-8 py-3 rounded-xl font-semibold hover:bg-green-700 transition disabled:opacity-50">
              {{ loading ? 'Menyimpan...' : (isEdit ? 'Update Produk' : 'Simpan Produk') }}
            </button>
            <RouterLink to="/admin/products"
              class="border border-gray-200 px-8 py-3 rounded-xl font-semibold text-gray-600 hover:bg-gray-50 transition">
              Batal
            </RouterLink>
          </div>
        </div>

        <!-- Sidebar Info -->
        <div class="space-y-6">
          <div class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100">
            <h3 class="font-bold text-gray-900 mb-4">Tips Gambar</h3>
            <ul class="text-sm text-gray-500 space-y-2">
              <li>✅ Format: JPG, PNG, WebP</li>
              <li>✅ Resolusi minimal 800×600px</li>
              <li>✅ Bisa upload banyak gambar sekaligus</li>
              <li>✅ Seret & lepas ke area upload</li>
              <li>⚠️ Maks. ukuran file: 5MB</li>
            </ul>
          </div>

          <div class="bg-green-50 rounded-2xl p-6 border border-green-100">
            <h3 class="font-bold text-green-800 mb-2">Status Upload</h3>
            <p class="text-xs text-green-700">
              Gambar diupload otomatis saat dipilih.<br/>
              Tunggu hingga semua selesai sebelum menyimpan produk.
            </p>
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
const uploadError = ref('')
const isDragging = ref(false)
const fileInput = ref(null)

const form = ref({ name: '', description: '' })
// imageList: [{ url: '/uploads/products/xxx.jpg', uploading: false }]
const imageList = ref([])

// Base URL backend — sesuaikan jika beda
const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

const getImageSrc = (url) => {
  if (!url) return ''
  // Jika sudah full URL, pakai langsung; jika path relatif, tambah base
  if (url.startsWith('http')) return url
  return BASE_URL + url
}

onMounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }

  if (isEdit.value) {
    const res = await api.get(`/products/${route.params.id}`)
    const data = res.data.data
    form.value = { name: data.name, description: data.description }
    if (data.Images) {
      imageList.value = data.Images.map(img => ({ url: img.image_url, uploading: false }))
    }
  }
})

const triggerFileInput = () => {
  fileInput.value?.click()
}

const handleFileChange = (e) => {
  const files = Array.from(e.target.files)
  uploadFiles(files)
  e.target.value = '' // reset agar bisa pilih file sama lagi
}

const handleDrop = (e) => {
  isDragging.value = false
  const files = Array.from(e.dataTransfer.files).filter(f =>
    ['image/jpeg', 'image/png', 'image/webp'].includes(f.type)
  )
  if (files.length === 0) {
    uploadError.value = 'Format file tidak didukung. Gunakan JPG, PNG, atau WebP.'
    return
  }
  uploadFiles(files)
}

const uploadFiles = async (files) => {
  uploadError.value = ''

  for (const file of files) {
    if (file.size > 5 * 1024 * 1024) {
      uploadError.value = `File "${file.name}" melebihi batas 5MB.`
      continue
    }

    // Tambah placeholder dengan status uploading
    const index = imageList.value.length
    imageList.value.push({ url: URL.createObjectURL(file), uploading: true })

    try {
      const formData = new FormData()
      formData.append('image', file)

      const res = await api.post('/admin/upload/image', formData, {
        headers: { 'Content-Type': 'multipart/form-data' }
      })

      // Ganti placeholder URL dengan URL dari server
      imageList.value[index] = { url: res.data.url, uploading: false }
    } catch (err) {
      imageList.value.splice(index, 1)
      uploadError.value = `Gagal mengupload "${file.name}". Coba lagi.`
    }
  }
}

const removeImage = (index) => {
  imageList.value.splice(index, 1)
}

const handleSubmit = async () => {
  if (!form.value.name) {
    error.value = 'Nama produk wajib diisi!'
    return
  }
  if (imageList.value.some(i => i.uploading)) {
    error.value = 'Tunggu hingga semua gambar selesai diupload.'
    return
  }

  loading.value = true
  error.value = ''

  try {
    const payload = {
      name: form.value.name,
      description: form.value.description,
      Images: imageList.value.map(img => ({ image_url: img.url }))
    }

    if (isEdit.value) {
      await api.put(`/admin/products/${route.params.id}`, payload)
    } else {
      await api.post('/admin/products', payload)
    }
    router.push('/admin/products')
  } catch (err) {
    error.value = 'Gagal menyimpan produk. Coba lagi.'
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