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
                <label class="text-sm font-semibold text-gray-700 block mb-2">Ringkasan</label>
                <textarea v-model="form.summary" rows="3" placeholder="Ringkasan singkat produk (opsional, tampil di halaman utama)"
                  class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500"></textarea>
              </div>
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-2">Deskripsi</label>
                <RichTextEditor v-model="form.description" placeholder="Masukkan deskripsi produk..." />
              </div>
              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="text-sm font-semibold text-gray-700 block mb-2">Kategori</label>
                  <input v-model="form.category" type="text" placeholder="Contoh: Pulp"
                    class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
                </div>
                <div>
                  <label class="text-sm font-semibold text-gray-700 block mb-2">Tags (koma)</label>
                  <input v-model="form.tags" type="text" placeholder="Contoh: BFK, High Quality, Export"
                    class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
                </div>
              </div>
              <div>
                <label class="flex items-center gap-2 text-sm text-gray-700">
                  <input type="checkbox" v-model="form.is_active"
                    class="rounded border-gray-300 text-green-600 focus:ring-green-500" />
                  Aktif (tampil di halaman publik)
                </label>
              </div>
            </div>
          </div>

          <!-- Image Upload -->
          <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100">
            <h3 class="font-bold text-gray-900 mb-2">Gambar Produk</h3>
            <p class="text-gray-400 text-xs mb-6">Upload gambar dari komputer. Bisa lebih dari satu gambar.</p>

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
                <div class="absolute bottom-0 left-0 right-0 bg-black/40 text-white text-xs px-2 py-1 text-center">
                  Gambar {{ i + 1 }}
                </div>
              </div>
            </div>

            <!-- Upload Area -->
            <div
              class="border-2 border-dashed border-gray-200 rounded-2xl p-8 text-center hover:border-green-400 transition cursor-pointer"
              @click="triggerFileInput"
              @dragover.prevent
              @drop.prevent="handleDrop">
              <input
                ref="fileInput"
                type="file"
                accept="image/*"
                multiple
                class="hidden"
                @change="handleFileChange"
              />
              <div v-if="uploading" class="flex flex-col items-center gap-3">
                <div class="w-8 h-8 border-4 border-green-600 border-t-transparent rounded-full animate-spin"></div>
                <p class="text-sm text-gray-500">Mengupload gambar...</p>
              </div>
              <div v-else class="flex flex-col items-center gap-3">
                <div class="w-14 h-14 bg-green-50 rounded-2xl flex items-center justify-center text-3xl">🖼️</div>
                <div>
                  <p class="font-semibold text-gray-700 text-sm">Klik untuk upload atau drag & drop</p>
                  <p class="text-gray-400 text-xs mt-1">PNG, JPG, WebP — Maksimal 5MB per file</p>
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
            <h3 class="font-bold text-gray-900 mb-4">📋 Panduan</h3>
            <ul class="text-sm text-gray-500 space-y-3">
              <li class="flex items-start gap-2">
                <span class="text-green-500 mt-0.5">✓</span>
                <span>Isi nama dan deskripsi produk</span>
              </li>
              <li class="flex items-start gap-2">
                <span class="text-green-500 mt-0.5">✓</span>
                <span>Upload gambar langsung dari komputer</span>
              </li>
              <li class="flex items-start gap-2">
                <span class="text-green-500 mt-0.5">✓</span>
                <span>Bisa upload lebih dari 1 gambar</span>
              </li>
              <li class="flex items-start gap-2">
                <span class="text-green-500 mt-0.5">✓</span>
                <span>Format: JPG, PNG, WebP</span>
              </li>
              <li class="flex items-start gap-2">
                <span class="text-yellow-500 mt-0.5">⚠</span>
                <span>Maksimal 5MB per gambar</span>
              </li>
            </ul>
          </div>

          <!-- Preview count -->
          <div class="bg-green-50 rounded-2xl p-6 border border-green-100 text-center">
            <p class="text-4xl font-black text-green-600">{{ imageList.length }}</p>
            <p class="text-sm text-green-700 mt-1">Gambar siap diupload</p>
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
import RichTextEditor from '../../components/RichTextEditor.vue'

const router = useRouter()
const route = useRoute()
const isEdit = computed(() => !!route.params.id)
const loading = ref(false)
const uploading = ref(false)
const error = ref('')
const fileInput = ref(null)

const form = ref({ name: '', summary: '', description: '', category: '', tags: '', is_active: true })
const imageList = ref([])

onMounted(async () => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }

  if (isEdit.value) {
    const res = await api.get(`/products/${route.params.id}`)
    const data = res.data.data
    form.value = {
      name: data.name,
      summary: data.summary || '',
      description: data.description,
      category: data.category || '',
      tags: data.tags || '',
      is_active: data.is_active,
    }
    if (data.Images) {
      imageList.value = data.Images.map(img => ({
        url: img.image_url,
        id: img.ID
      }))
    }
  }
})

const triggerFileInput = () => {
  fileInput.value.click()
}

const handleFileChange = async (e) => {
  const files = Array.from(e.target.files)
  await uploadFiles(files)
}

const handleDrop = async (e) => {
  const files = Array.from(e.dataTransfer.files).filter(f => f.type.startsWith('image/'))
  await uploadFiles(files)
}

const uploadFiles = async (files) => {
  if (files.length === 0) {return}
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
    } catch (_err) {
      error.value = `Gagal upload ${file.name}`
    }
  }

  uploading.value = false
  // Reset input
  if (fileInput.value) {fileInput.value.value = ''}
}

const removeImage = (index) => {
  imageList.value.splice(index, 1)
}

const handleSubmit = async () => {
  if (!form.value.name) {
    error.value = 'Nama produk wajib diisi!'
    return
  }
  loading.value = true
  error.value = ''

  try {
    const payload = {
      name: form.value.name,
      summary: form.value.summary,
      description: form.value.description,
      category: form.value.category,
      tags: form.value.tags,
      is_active: form.value.is_active,
      Images: imageList.value.map(img => ({ image_url: img.url }))
    }

    if (isEdit.value) {
      await api.put(`/admin/products/${route.params.id}`, payload)
    } else {
      await api.post('/admin/products', payload)
    }
    router.push('/admin/products')
  } catch (_err) {
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