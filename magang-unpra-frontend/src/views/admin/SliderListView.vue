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
        <RouterLink to="/admin/slider"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
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
      <div class="flex justify-between items-center mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Product Slider</h1>
        </div>
        <button @click="openCreate"
          class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-2">
          <span class="text-lg leading-none">+</span> Tambah Slide
        </button>
      </div>

      <div v-if="alert.show"
        :class="['mb-4 px-4 py-3 rounded-xl text-sm font-medium flex items-center gap-2',
          alert.type === 'success' ? 'bg-green-50 text-green-700 border border-green-200' : 'bg-red-50 text-red-700 border border-red-200']">
        {{ alert.type === 'success' ? '✅' : '❌' }} {{ alert.message }}
      </div>

      <div v-if="loading" class="flex items-center justify-center py-20 gap-3 text-gray-400">
        <div class="w-5 h-5 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
        Memuat data...
      </div>

      <div v-else-if="slides.length === 0"
        class="bg-white rounded-2xl p-16 text-center border border-gray-100">
        <p class="text-4xl mb-4">🖼️</p>
        <p class="text-gray-500 mb-4">Belum ada slide.</p>
        <button @click="openCreate"
          class="bg-green-600 text-white px-6 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
          + Tambah Slide
        </button>
      </div>

      <div v-else class="grid gap-6">
        <div v-for="(slide, i) in slides" :key="slide.ID || i"
          class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100 flex items-center gap-6">
          <div class="w-28 h-20 rounded-xl overflow-hidden bg-gray-100 flex-shrink-0 border">
            <img :src="getImageUrl(slide.image_url)" class="w-full h-full object-cover"
              @error="(e) => e.target.src='https://via.placeholder.com/150?text=No+Image'" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="font-semibold text-gray-900">{{ slide.title || '(no title)' }}</p>
            <p class="text-sm text-gray-400 truncate">{{ slide.description }}</p>
            <p class="text-xs text-gray-300 mt-1">Urutan: {{ slide.sort_order || 0 }}</p>
          </div>
          <div class="flex gap-2 flex-shrink-0">
            <button @click="openEdit(slide)"
              class="bg-blue-50 text-blue-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-blue-100 transition">Edit</button>
            <button @click="confirmDelete(slide)"
              class="bg-red-50 text-red-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-red-100 transition">Hapus</button>
          </div>
        </div>
      </div>
    </main>

    <!-- Modal Form -->
    <Transition name="modal">
      <div v-if="showModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="closeModal">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-lg max-h-[90vh] overflow-y-auto">
          <div class="flex items-center justify-between p-6 border-b border-gray-100">
            <h2 class="text-xl font-black text-gray-900">
              {{ isEdit ? '✏️ Edit Slide' : '➕ Tambah Slide' }}
            </h2>
            <button @click="closeModal" class="text-gray-400 hover:text-gray-600 text-2xl leading-none">×</button>
          </div>
          <form @submit.prevent="submitForm" class="p-6 space-y-5">
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Judul <span class="text-red-500">*</span></label>
              <input v-model="form.title" type="text" placeholder="Judul slide..." required
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400" />
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Deskripsi</label>
              <textarea v-model="form.description" rows="2" placeholder="Deskripsi slide..."
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400 resize-none"></textarea>
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Gambar</label>
              <div class="border-2 border-dashed border-gray-200 rounded-xl p-4 text-center hover:border-green-400 transition cursor-pointer"
                @click="$refs.fileInput.click()">
                <input ref="fileInput" type="file" accept="image/*" class="hidden" @change="handleFileChange" />
                <div v-if="imagePreview || form.image_url">
                  <img :src="imagePreview || getImageUrl(form.image_url)"
                    class="mx-auto h-28 object-cover rounded-lg mb-2" />
                  <p class="text-xs text-gray-400">Klik untuk ganti gambar</p>
                </div>
                <div v-else>
                  <p class="text-3xl mb-2">🖼️</p>
                  <p class="text-sm text-gray-500">Klik untuk upload gambar</p>
                </div>
              </div>
            </div>
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Urutan</label>
              <input v-model.number="form.sort_order" type="number" min="0"
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400" />
            </div>
            <div v-if="formError" class="bg-red-50 text-red-600 text-sm px-4 py-3 rounded-xl border border-red-200">
              ❌ {{ formError }}
            </div>
            <div class="flex gap-3 pt-2">
              <button type="button" @click="closeModal"
                class="flex-1 border border-gray-200 text-gray-600 py-2.5 rounded-xl text-sm font-semibold hover:bg-gray-50 transition">Batal</button>
              <button type="submit" :disabled="submitting"
                class="flex-1 bg-green-600 text-white py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition disabled:opacity-60 flex items-center justify-center gap-2">
                <div v-if="submitting" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                {{ submitting ? 'Menyimpan...' : (isEdit ? 'Simpan Perubahan' : 'Tambah Slide') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Transition>

    <!-- Modal Hapus -->
    <Transition name="modal">
      <div v-if="showDeleteModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="showDeleteModal = false">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-sm p-6 text-center">
          <p class="text-4xl mb-3">🗑️</p>
          <h3 class="text-lg font-black text-gray-900 mb-1">Hapus Slide?</h3>
          <p class="text-sm text-gray-500 mb-6">Slide akan dihapus permanen.</p>
          <div class="flex gap-3">
            <button @click="showDeleteModal = false"
              class="flex-1 border border-gray-200 text-gray-600 py-2.5 rounded-xl text-sm font-semibold hover:bg-gray-50 transition">Batal</button>
            <button @click="doDelete" :disabled="submitting"
              class="flex-1 bg-red-600 text-white py-2.5 rounded-xl text-sm font-semibold hover:bg-red-700 transition disabled:opacity-60">
              {{ submitting ? 'Menghapus...' : 'Hapus' }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const BASE_URL = 'http://localhost:8080'
const router = useRouter()
const slides = ref([])
const loading = ref(true)
const submitting = ref(false)
const showModal = ref(false)
const showDeleteModal = ref(false)
const isEdit = ref(false)
const selectedSlide = ref(null)
const imageFile = ref(null)
const imagePreview = ref(null)
const formError = ref('')
const alert = ref({ show: false, type: 'success', message: '' })

const form = ref({ title: '', description: '', image_url: '', sort_order: 0 })

const getImageUrl = (path) => {
  if (!path) {return ''}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const showAlert = (type, message) => {
  alert.value = { show: true, type, message }
  setTimeout(() => { alert.value.show = false }, 3000)
}

const fetchSlides = async () => {
  loading.value = true
  try {
    const res = await api.get('/admin/product-slides')
    slides.value = res.data?.data || res.data || []
  } catch { showAlert('error', 'Gagal memuat slide') }
  finally { loading.value = false }
}

const openCreate = () => {
  isEdit.value = false
  form.value = { title: '', description: '', image_url: '', sort_order: 0 }
  imagePreview.value = null; imageFile.value = null; formError.value = ''
  showModal.value = true
}

const openEdit = (slide) => {
  isEdit.value = true; selectedSlide.value = slide
  form.value = {
    title: slide.title || '',
    description: slide.description || '',
    image_url: slide.image_url || '',
    sort_order: slide.sort_order || 0,
  }
  imagePreview.value = null; imageFile.value = null; formError.value = ''
  showModal.value = true
}

const closeModal = () => { showModal.value = false; formError.value = '' }

const handleFileChange = async (e) => {
  const file = e.target.files[0]
  if (!file) {return}
  if (file.size > 5 * 1024 * 1024) { formError.value = 'Ukuran gambar maksimal 5MB'; return }
  imageFile.value = file
  imagePreview.value = URL.createObjectURL(file)
  await uploadImage()
}

const uploadImage = async () => {
  if (!imageFile.value) {return}
  const fd = new FormData()
  fd.append('image', imageFile.value)
  try {
    const res = await api.post('/upload', fd, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
    form.value.image_url = res.data?.url || ''
  } catch (e) {
    console.error('Upload error:', e)
    formError.value = 'Gagal upload gambar'
  }
}

const submitForm = async () => {
  if (!form.value.title.trim()) {
    formError.value = 'Judul wajib diisi'; return
  }
  submitting.value = true; formError.value = ''
  try {
    const payload = {
      title: form.value.title,
      description: form.value.description,
      image_url: form.value.image_url,
      sort_order: form.value.sort_order,
    }
    if (isEdit.value) {
      await api.put(`/admin/product-slides/${selectedSlide.value.ID}`, payload)
      showAlert('success', 'Slide berhasil diperbarui!')
    } else {
      await api.post('/admin/product-slides', payload)
      showAlert('success', 'Slide berhasil ditambahkan!')
    }
    closeModal(); await fetchSlides()
  } catch (e) {
    console.error('Submit error:', e)
    formError.value = e.response?.data?.error || e.response?.data?.message || e.message || 'Terjadi kesalahan. Coba lagi.'
  } finally { submitting.value = false }
}

const confirmDelete = (slide) => { selectedSlide.value = slide; showDeleteModal.value = true }
const doDelete = async () => {
  submitting.value = true
  try {
    await api.delete(`/admin/product-slides/${selectedSlide.value.ID}`)
    slides.value = slides.value.filter(s => s.ID !== selectedSlide.value.ID)
    showDeleteModal.value = false
    showAlert('success', 'Slide berhasil dihapus!')
  } catch { showAlert('error', 'Gagal menghapus slide') }
  finally { submitting.value = false }
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}

onMounted(() => {
  if (!localStorage.getItem('token')) { router.push('/admin/login'); return }
  fetchSlides()
})
</script>

<style scoped>
.modal-enter-active, .modal-leave-active { transition: opacity 0.2s ease; }
.modal-enter-from, .modal-leave-to { opacity: 0; }
</style>
