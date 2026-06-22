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
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl hover:bg-gray-800 text-sm text-gray-300 transition">
          🏢 About Section
        </RouterLink>
        <RouterLink to="/admin/community"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl bg-green-600 text-white text-sm font-medium">
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
      <div class="mb-8 flex items-center justify-between">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Local Community Development</h1>
          <p class="text-gray-400 text-sm mt-1">Kelola kartu layanan masyarakat yang tampil di halaman utama.</p>
        </div>
        <button @click="openCreate"
          class="bg-green-600 text-white px-5 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-2">
          + Tambah Kartu
        </button>
      </div>

      <div v-if="alert.show"
        :class="['mb-4 px-4 py-3 rounded-xl text-sm font-medium flex items-center gap-2',
          alert.type === 'success' ? 'bg-green-50 text-green-700 border border-green-200' : 'bg-red-50 text-red-700 border border-red-200']">
        {{ alert.type === 'success' ? '✅' : '❌' }} {{ alert.message }}
      </div>

      <div v-if="loading" class="flex items-center justify-center py-20 text-gray-400">
        <div class="w-5 h-5 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
      </div>

      <div v-else-if="!cards.length" class="bg-white rounded-2xl p-12 shadow-sm border border-gray-100 text-center">
        <p class="text-gray-400 text-lg">Belum ada kartu community. Klik "Tambah Kartu" untuk memulai.</p>
      </div>

      <div v-else class="grid gap-4">
        <div v-for="card in cards" :key="card.ID"
          class="bg-white rounded-2xl p-6 shadow-sm border border-gray-100 flex items-start gap-5">
          <div class="flex flex-col gap-1">
            <div v-if="card.Images && card.Images.length" class="flex -space-x-2">
              <div v-for="(img, i) in card.Images.slice(0, 3)" :key="img.ID"
                class="w-10 h-10 rounded-full border-2 border-white overflow-hidden bg-gray-100"
                :class="{ 'z-30': i === 0, 'z-20': i === 1, 'z-10': i === 2 }">
                <img :src="getImageUrl(img.image_url)" class="w-full h-full object-cover" />
              </div>
              <div v-if="card.Images.length > 3"
                class="w-10 h-10 rounded-full border-2 border-white bg-gray-200 flex items-center justify-center text-xs font-bold text-gray-500 z-0">
                +{{ card.Images.length - 3 }}
              </div>
            </div>
            <div v-else class="w-10 h-10 rounded-xl bg-gray-100 flex items-center justify-center text-lg text-gray-300">
              🏷️
            </div>
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-3 mb-1">
              <h3 class="font-bold text-gray-900">{{ card.title }}</h3>
              <span v-if="!card.is_active"
                class="text-xs bg-red-100 text-red-600 px-2 py-0.5 rounded-full font-semibold">Nonaktif</span>
            </div>
            <p class="text-sm text-gray-500 line-clamp-2">{{ card.description }}</p>
            <p class="text-xs text-gray-400 mt-1">Urutan: {{ card.sort_order }} | Gambar: {{ card.Images?.length || 0 }}</p>
          </div>
          <div class="flex gap-2 flex-shrink-0">
            <button @click="openEdit(card)"
              class="px-4 py-2 text-sm font-semibold text-gray-600 border border-gray-200 rounded-xl hover:bg-gray-50 transition">
              Edit
            </button>
            <button @click="deleteCard(card.ID)"
              class="px-4 py-2 text-sm font-semibold text-red-600 border border-red-200 rounded-xl hover:bg-red-50 transition">
              Hapus
            </button>
          </div>
        </div>
      </div>

      <!-- Modal Create/Edit -->
      <Transition name="modal-fade">
        <div v-if="showModal" class="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm p-4"
          @click.self="showModal = false">
          <div class="bg-white rounded-2xl w-full max-w-lg p-8 shadow-2xl">
            <h2 class="text-xl font-black text-gray-900 mb-6">{{ editing ? 'Edit Kartu' : 'Tambah Kartu' }}</h2>

            <div class="space-y-4">
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-1.5">Judul</label>
                <input v-model="form.title" type="text" placeholder="Contoh: Education"
                  class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
              </div>
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-1.5">Deskripsi</label>
                <textarea v-model="form.description" rows="4" placeholder="Deskripsi singkat..."
                  class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-500 resize-none"></textarea>
              </div>
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-1.5">Ikon Utama</label>
                <div class="flex items-center gap-4">
                  <div v-if="form.icon_path" class="w-16 h-16 rounded-xl overflow-hidden bg-gray-100 border border-gray-200 flex-shrink-0">
                    <img :src="getImageUrl(form.icon_path)" class="w-full h-full object-cover" />
                  </div>
                  <div class="flex-1">
                    <input type="file" accept="image/*" ref="fileInput" class="hidden" @change="onIconChange" />
                    <button @click="$refs.fileInput.click()"
                      class="border border-dashed border-gray-300 rounded-xl px-4 py-3 text-sm text-gray-500 hover:border-green-500 hover:text-green-600 transition w-full text-left">
                      {{ uploadingIcon ? 'Mengupload...' : 'Pilih ikon dari komputer' }}
                    </button>
                    <p class="text-xs text-gray-400 mt-1">JPG, PNG, WebP. Maks 5MB.</p>
                  </div>
                </div>
              </div>

              <!-- Multiple Images -->
              <div>
                <label class="text-sm font-semibold text-gray-700 block mb-1.5">Gambar Tambahan</label>
                <div class="flex flex-wrap gap-2 mb-3">
                  <div v-for="(img, i) in imageList" :key="i" class="relative group">
                    <img :src="getImageUrl(img.url)" class="w-20 h-20 rounded-xl object-cover border border-gray-200" />
                    <button @click="imageList.splice(i, 1)"
                      class="absolute -top-2 -right-2 w-5 h-5 bg-red-500 text-white rounded-full text-xs flex items-center justify-center opacity-0 group-hover:opacity-100 transition">✕</button>
                  </div>
                  <div v-if="uploadingImages" class="w-20 h-20 rounded-xl border-2 border-dashed border-gray-300 flex items-center justify-center text-gray-400 text-sm">
                    <div class="w-5 h-5 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
                  </div>
                  <label class="w-20 h-20 rounded-xl border-2 border-dashed border-gray-300 flex items-center justify-center cursor-pointer hover:border-green-500 hover:text-green-600 transition text-gray-400 text-2xl">
                    +
                    <input type="file" accept="image/*" class="hidden" multiple @change="onImagesChange" />
                  </label>
                </div>
                <p class="text-xs text-gray-400">Upload beberapa gambar untuk kartu ini.</p>
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="text-sm font-semibold text-gray-700 block mb-1.5">Link Tujuan</label>
                  <input v-model="form.link" type="text" placeholder="/sustainability/csr/community"
                    class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
                </div>
                <div>
                  <label class="text-sm font-semibold text-gray-700 block mb-1.5">Urutan Tampil</label>
                  <input v-model.number="form.sort_order" type="number" min="0"
                    class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-500" />
                </div>
              </div>
              <div>
                <label class="flex items-center gap-3 cursor-pointer">
                  <input type="checkbox" v-model="form.is_active"
                    class="w-5 h-5 rounded-lg border-gray-300 text-green-600 focus:ring-green-500" />
                  <span class="text-sm text-gray-700">Aktif</span>
                </label>
              </div>
            </div>

            <div class="flex gap-3 mt-6 pt-4 border-t border-gray-100">
              <button @click="save" :disabled="saving"
                class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition disabled:opacity-50 flex items-center gap-2">
                <div v-if="saving" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                {{ saving ? 'Menyimpan...' : 'Simpan' }}
              </button>
              <button @click="showModal = false"
                class="border border-gray-200 px-6 py-2.5 rounded-xl text-sm font-semibold text-gray-600 hover:bg-gray-50 transition">
                Batal
              </button>
            </div>
          </div>
        </div>
      </Transition>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const BASE_URL = import.meta.env.VITE_BASE_URL || 'http://localhost:8080'
const router = useRouter()
const loading = ref(true)
const saving = ref(false)
const uploadingIcon = ref(false)
const uploadingImages = ref(false)
const showModal = ref(false)
const editing = ref(false)
const editingId = ref(null)
const cards = ref([])
const imageList = ref([])
const alert = ref({ show: false, type: 'success', message: '' })
const fileInput = ref(null)

const form = ref({
  title: '',
  description: '',
  icon_path: '',
  link: '/sustainability/csr/community',
  sort_order: 0,
  is_active: true
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

const fetchCards = async () => {
  loading.value = true
  try {
    const res = await api.get('/admin/community-cards')
    cards.value = res.data?.data || []
  } catch { showAlert('error', 'Gagal memuat data') }
  finally { loading.value = false }
}

const resetForm = () => {
  form.value = { title: '', description: '', icon_path: '', link: '/sustainability/csr/community', sort_order: 0, is_active: true }
  imageList.value = []
}

const openCreate = () => {
  editing.value = false
  editingId.value = null
  resetForm()
  showModal.value = true
}

const openEdit = (card) => {
  editing.value = true
  editingId.value = card.ID
  form.value = {
    title: card.title || '',
    description: card.description || '',
    icon_path: card.icon_path || '',
    link: card.link || '/sustainability/csr/community',
    sort_order: card.sort_order || 0,
    is_active: card.is_active !== false
  }
  imageList.value = (card.Images || []).map(img => ({ url: img.image_url }))
  showModal.value = true
}

const uploadFile = async (file) => {
  const fd = new FormData()
  fd.append('image', file)
  const res = await api.post('/upload', fd, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
  return `uploads/${res.data?.filename || ''}`
}

const onIconChange = async (e) => {
  const file = e.target.files[0]
  if (!file) {return}
  uploadingIcon.value = true
  try {
    form.value.icon_path = await uploadFile(file)
    showAlert('success', 'Ikon berhasil diupload!')
  } catch {
    showAlert('error', 'Gagal upload ikon')
  } finally {
    uploadingIcon.value = false
    if (fileInput.value) {fileInput.value.value = ''}
  }
}

const onImagesChange = async (e) => {
  const files = Array.from(e.target.files)
  if (!files.length) {return}
  uploadingImages.value = true
  for (const file of files) {
    if (file.size > 5 * 1024 * 1024) {
      showAlert('error', `File ${file.name} terlalu besar (maks 5MB)`)
      continue
    }
    try {
      const url = await uploadFile(file)
      imageList.value.push({ url })
    } catch {
      showAlert('error', `Gagal upload ${file.name}`)
    }
  }
  uploadingImages.value = false
  e.target.value = ''
}

const save = async () => {
  saving.value = true
  try {
    const payload = {
      ...form.value,
      Images: imageList.value.map(img => ({ image_url: img.url }))
    }
    if (editing.value) {
      await api.put(`/admin/community-cards/${editingId.value}`, payload)
    } else {
      await api.post('/admin/community-cards', payload)
    }
    showAlert('success', 'Kartu berhasil disimpan!')
    showModal.value = false
    await fetchCards()
  } catch (e) {
    showAlert('error', e.response?.data?.error || e.message || 'Gagal menyimpan')
  } finally { saving.value = false }
}

const deleteCard = async (id) => {
  if (!confirm('Yakin ingin menghapus kartu ini?')) {return}
  try {
    await api.delete(`/admin/community-cards/${id}`)
    showAlert('success', 'Kartu berhasil dihapus!')
    await fetchCards()
  } catch {
    showAlert('error', 'Gagal menghapus kartu')
  }
}

const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}

onMounted(() => {
  if (!localStorage.getItem('token')) { router.push('/admin/login'); return }
  fetchCards()
})
</script>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.2s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
</style>
