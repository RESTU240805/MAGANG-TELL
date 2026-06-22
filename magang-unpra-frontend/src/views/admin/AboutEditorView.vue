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
        <RouterLink to="/admin/product-page"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          📦 Product
        </RouterLink>
        <RouterLink to="/admin/about"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
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
      <div class="mb-8">
        <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
        <h1 class="text-3xl font-black text-gray-900">About Section</h1>
        <p class="text-gray-400 text-sm mt-1">Kelola konten "About TELPP" yang tampil di halaman utama.</p>
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
            <label class="text-sm font-semibold text-gray-700 block mb-2">Judul Section</label>
            <input v-model="form.title" type="text" placeholder="ABOUT TELPP"
              class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
          </div>

          <div>
            <label class="text-sm font-semibold text-gray-700 block mb-2">Deskripsi</label>
            <RichTextEditor v-model="form.description" placeholder="Deskripsi tentang perusahaan..." />
          </div>

          <div>
            <label class="text-sm font-semibold text-gray-700 block mb-2">Gambar</label>
            <div class="flex items-center gap-4">
              <div v-if="form.image_path" class="w-40 h-28 rounded-xl overflow-hidden bg-gray-100 border border-gray-200">
                <img :src="getImageUrl(form.image_path)" class="w-full h-full object-cover" />
              </div>
              <div class="flex-1">
                <input type="file" accept="image/*" @change="onImageChange" ref="fileInput"
                  class="hidden" />
                <button @click="$refs.fileInput.click()"
                  class="border border-dashed border-gray-300 rounded-xl px-6 py-4 text-sm text-gray-500 hover:border-green-500 hover:text-green-600 transition w-full text-left">
                  {{ uploading ? 'Mengupload...' : 'Klik untuk upload gambar baru' }}
                </button>
                <p class="text-xs text-gray-400 mt-1">Format: JPG, PNG. Maks 5MB.</p>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="text-sm font-semibold text-gray-700 block mb-2">Badge Number</label>
              <input v-model="form.badge_number" type="text" placeholder="20+"
                class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
            <div>
              <label class="text-sm font-semibold text-gray-700 block mb-2">Badge Label</label>
              <input v-model="form.badge_label" type="text" placeholder="Years of Excellence"
                class="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
            </div>
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
      <div class="mt-8 max-w-3xl">
        <h3 class="text-sm font-semibold text-gray-500 tracking-widest mb-3">PRATINJAU</h3>
        <div class="bg-white rounded-2xl p-8 shadow-sm border border-gray-100">
          <div class="flex gap-8 items-center">
            <div class="flex-1">
              <p class="text-xs font-bold tracking-widest text-green-600 mb-2 uppercase">About TELPP</p>
              <h2 class="text-xl font-black text-gray-900 mb-3">{{ form.title || 'ABOUT TELPP' }}</h2>
              <p class="text-sm text-gray-500 leading-relaxed" v-html="form.description || 'Deskripsi akan muncul di sini...'"></p>
            </div>
            <div class="w-64 h-44 rounded-xl overflow-hidden bg-gray-100 relative flex-shrink-0">
              <img v-if="form.image_path" :src="getImageUrl(form.image_path)" class="w-full h-full object-cover" />
              <div v-if="form.badge_number" class="absolute bottom-0 right-0 bg-green-700 text-white px-4 py-3 text-center rounded-tl-xl">
                <span class="block text-xl font-black">{{ form.badge_number }}</span>
                <span class="block text-[9px] font-bold tracking-widest uppercase">{{ form.badge_label }}</span>
              </div>
            </div>
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
import RichTextEditor from '../../components/RichTextEditor.vue'

const BASE_URL = 'http://localhost:8080'
const router = useRouter()
const loading = ref(true)
const saving = ref(false)
const uploading = ref(false)
const alert = ref({ show: false, type: 'success', message: '' })
const fileInput = ref(null)

const form = ref({
  title: '',
  description: '',
  image_path: '',
  badge_number: '',
  badge_label: ''
})

const getImageUrl = (path) => {
  if (!path) {return ''}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const showAlert = (type, message) => {
  alert.value = { show: true, type, message }
  setTimeout(() => { alert.value.show = false }, 3000)
}

const fetchData = async () => {
  loading.value = true
  try {
    const res = await api.get('/about-section')
    const d = res.data?.data
    if (d) {
      form.value = {
        title: d.title || '',
        description: d.description || '',
        image_path: d.image_path || '',
        badge_number: d.badge_number || '',
        badge_label: d.badge_label || ''
      }
    }
  } catch { showAlert('error', 'Gagal memuat data') }
  finally { loading.value = false }
}

const onImageChange = async (e) => {
  const file = e.target.files[0]
  if (!file) {return}
  uploading.value = true
  try {
    const fd = new FormData()
    fd.append('image', file)
    const res = await api.post('/upload', fd, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
    form.value.image_path = `uploads/${res.data?.filename || ''}`
    showAlert('success', 'Gambar berhasil diupload!')
  } catch (_err) {
    showAlert('error', 'Gagal upload gambar')
  } finally { uploading.value = false }
}

const save = async () => {
  saving.value = true
  try {
    await api.put('/about-section', form.value)
    showAlert('success', 'About section berhasil disimpan!')
  } catch (e) {
    showAlert('error', e.response?.data?.error || e.message || 'Gagal menyimpan')
  } finally { saving.value = false }
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}

onMounted(() => {
  if (!localStorage.getItem('token')) { router.push('/admin/login'); return }
  fetchData()
})
</script>
