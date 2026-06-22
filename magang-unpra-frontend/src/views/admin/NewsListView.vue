<template>
  <div class="min-h-screen bg-gray-100 flex">

    <!-- ─── Sidebar ─────────────────────────────────────────── -->
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

    <!-- ─── Main Content ────────────────────────────────────── -->
    <main class="flex-1 ml-64 p-10">

      <!-- Header -->
      <div class="flex justify-between items-center mb-8">
        <div>
          <p class="text-xs text-gray-400 font-semibold tracking-widest">CONTENT ENGINE</p>
          <h1 class="text-3xl font-black text-gray-900">Corporate News</h1>
        </div>
        <!-- Tombol Tambah — tidak tertutup navbar karena main pakai ml-64 -->
        <button
          @click="openCreate"
          class="bg-green-600 text-white px-6 py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition flex items-center gap-2">
          <span class="text-lg leading-none">+</span> Tambah Berita
        </button>
      </div>

      <!-- Alert notifikasi -->
      <div v-if="alert.show"
        :class="['mb-4 px-4 py-3 rounded-xl text-sm font-medium flex items-center gap-2',
          alert.type === 'success' ? 'bg-green-50 text-green-700 border border-green-200' : 'bg-red-50 text-red-700 border border-red-200']">
        {{ alert.type === 'success' ? '✅' : '❌' }} {{ alert.message }}
      </div>

      <!-- Loading -->
      <div v-if="loading" class="flex items-center justify-center py-20 gap-3 text-gray-400">
        <div class="w-5 h-5 border-2 border-gray-200 border-t-green-500 rounded-full animate-spin"></div>
        Memuat data...
      </div>

      <!-- Empty -->
      <div v-else-if="newsList.length === 0"
        class="bg-white rounded-2xl p-16 text-center border border-gray-100">
        <p class="text-4xl mb-4">📰</p>
        <p class="text-gray-500 mb-4">Belum ada berita. Tambahkan berita pertama!</p>
        <button @click="openCreate"
          class="bg-green-600 text-white px-6 py-2 rounded-xl text-sm font-semibold hover:bg-green-700 transition">
          + Tambah Berita
        </button>
      </div>

      <!-- Tabel Berita -->
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
                <div class="flex items-center gap-3">
                  <img v-if="news.thumbnail_path"
                    :src="getImageUrl(news.thumbnail_path)"
                    class="w-12 h-12 rounded-lg object-cover flex-shrink-0 bg-gray-100"
                    @error="(e) => e.target.style.display='none'"
                  />
                  <div class="w-10 h-12 rounded-lg bg-gray-100 flex items-center justify-center flex-shrink-0 text-lg" v-else>📰</div>
                  <div>
                    <div class="flex items-center gap-2 mb-0.5">
                      <p class="font-semibold text-gray-900 text-sm">{{ news.title }}</p>
                      <span v-if="news.is_published"
                        class="bg-green-100 text-green-600 text-[10px] font-bold px-1.5 py-0.5 rounded-full">LIVE</span>
                      <span v-else
                        class="bg-gray-100 text-gray-400 text-[10px] font-bold px-1.5 py-0.5 rounded-full">DRAFT</span>
                    </div>
                    <p class="text-gray-400 text-xs line-clamp-1">{{ news.summary || news.content }}</p>
                  </div>
                </div>
              </td>
              <td class="px-6 py-4">
                <span class="bg-green-100 text-green-700 px-3 py-1 rounded-full text-xs font-semibold">
                  {{ news.category || 'Umum' }}
                </span>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">
                {{ formatDate(news.published_at || news.CreatedAt) }}
              </td>
              <td class="px-6 py-4">
                <div class="flex gap-2">
                  <button @click="openEdit(news)"
                    class="bg-blue-50 text-blue-600 px-3 py-1.5 rounded-lg text-xs font-semibold hover:bg-blue-100 transition">
                    Edit
                  </button>
                  <button @click="confirmDelete(news)"
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

    <!-- ─── Modal Form (Create / Edit) ─────────────────────── -->
    <Transition name="modal">
      <div v-if="showModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="closeModal">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">

          <!-- Modal Header -->
          <div class="flex items-center justify-between p-6 border-b border-gray-100">
            <h2 class="text-xl font-black text-gray-900">
              {{ isEdit ? '✏️ Edit Berita' : '➕ Tambah Berita' }}
            </h2>
            <button @click="closeModal" class="text-gray-400 hover:text-gray-600 text-2xl leading-none">×</button>
          </div>

          <!-- Modal Body -->
          <form @submit.prevent="submitForm" class="p-6 space-y-5">

            <!-- Judul -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Judul Berita <span class="text-red-500">*</span></label>
              <input
                v-model="form.title"
                type="text"
                placeholder="Masukkan judul berita..."
                required
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent"
              />
            </div>

            <!-- Kategori -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Kategori</label>
              <div class="flex gap-2 flex-wrap mb-2">
                <button
                  v-for="cat in categories" :key="cat" type="button"
                  @click="form.category = cat"
                  :class="['px-3 py-1 rounded-full text-xs font-semibold border transition',
                    form.category === cat
                      ? 'bg-green-600 text-white border-green-600'
                      : 'bg-white text-gray-500 border-gray-200 hover:border-green-400']">
                  {{ cat }}
                </button>
              </div>
              <input
                v-model="form.category"
                type="text"
                placeholder="Atau ketik kategori sendiri..."
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent"
              />
            </div>

            <!-- Slug -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">
                Slug
                <span class="text-xs text-gray-400 font-normal ml-1">(otomatis dari judul)</span>
              </label>
              <input
                v-model="form.slug"
                type="text"
                placeholder="contoh-judul-berita"
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent font-mono"
              />
            </div>

            <!-- Summary -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Ringkasan</label>
              <textarea
                v-model="form.summary"
                rows="2"
                placeholder="Ringkasan singkat berita (tampil di card)..."
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent resize-none"
              ></textarea>
            </div>

            <!-- Konten -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Konten <span class="text-red-500">*</span></label>

              <!-- Toolbar -->
              <div class="flex items-center gap-1 border border-b-0 border-gray-200 rounded-t-xl px-2 py-1.5 bg-gray-50">
                <button type="button" @click="execCmd('bold')" title="Bold"
                  class="p-1.5 rounded hover:bg-gray-200 transition text-gray-600 font-bold text-sm">B</button>
                <button type="button" @click="execCmd('italic')" title="Italic"
                  class="p-1.5 rounded hover:bg-gray-200 transition text-gray-600 italic text-sm">I</button>
                <span class="w-px h-5 bg-gray-300 mx-1"></span>
                <button type="button" @click="execCmd('insertUnorderedList')" title="Bullet List"
                  class="p-1.5 rounded hover:bg-gray-200 transition text-gray-600 text-sm">•≡</button>
                <button type="button" @click="execCmd('insertOrderedList')" title="Numbered List"
                  class="p-1.5 rounded hover:bg-gray-200 transition text-gray-600 text-sm">1≡</button>
                <span class="w-px h-5 bg-gray-300 mx-1"></span>
                <button type="button" @click="execCmd('formatBlock', 'blockquote')" title="Quote"
                  class="p-1.5 rounded hover:bg-gray-200 transition text-gray-600 text-sm">❝</button>
                <span class="w-px h-5 bg-gray-300 mx-1"></span>
                <button type="button" @click="triggerContentImage" title="Insert Image"
                  class="p-1.5 rounded hover:bg-green-100 transition text-green-700 text-sm font-semibold">
                  🖼️ Gambar
                </button>
                <input ref="contentImageInput" type="file" accept="image/*" class="hidden" @change="insertContentImage" />
              </div>

              <!-- Editor -->
              <div
                ref="contentEditor"
                contenteditable="true"
                @input="onContentInput"
                class="w-full border border-gray-200 rounded-b-xl px-4 py-3 text-sm min-h-[160px] max-h-[400px] overflow-y-auto focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent prose prose-sm max-w-none [&_img]:max-w-full [&_img]:rounded-lg [&_img]:my-3 [&_img]:shadow-sm [&_ul]:list-disc [&_ul]:pl-6 [&_ol]:list-decimal [&_ol]:pl-6 [&_blockquote]:border-l-4 [&_blockquote]:border-green-500 [&_blockquote]:pl-4 [&_blockquote]:italic [&_blockquote]:text-gray-500"
                :data-placeholder="'Tulis isi berita di sini...'"
              ></div>
              <p v-if="contentUploading" class="text-xs text-gray-400 mt-1">Mengupload gambar...</p>
            </div>

            <!-- Tanggal -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Tanggal Berita</label>
              <input
                v-model="form.published_at"
                type="date"
                class="w-full border border-gray-200 rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-2 focus:ring-green-400 focus:border-transparent"
              />
            </div>

            <!-- Upload Gambar -->
            <div>
              <label class="block text-sm font-semibold text-gray-700 mb-1.5">Gambar</label>
              <div v-if="imageList.length" class="grid grid-cols-2 md:grid-cols-3 gap-3 mb-4">
                <div v-for="(img, i) in imageList" :key="i" class="relative rounded-xl overflow-hidden h-28 bg-gray-100 group">
                  <img :src="getImageUrl(img.image_url)" class="w-full h-full object-cover" />
                  <button type="button" @click="removeImage(i)"
                    class="absolute inset-0 bg-black/50 text-white text-xs font-semibold opacity-0 group-hover:opacity-100 transition">
                    Hapus
                  </button>
                </div>
              </div>
              <div
                class="border-2 border-dashed border-gray-200 rounded-xl p-4 text-center hover:border-green-400 transition cursor-pointer"
                @click="$refs.fileInput.click()"
                @dragover.prevent
                @drop.prevent="handleDrop">
                <input ref="fileInput" type="file" accept="image/*" multiple class="hidden" @change="handleFileChange" />
                <p class="text-3xl mb-2">🖼️</p>
                <p class="text-sm text-gray-500">Klik atau drag & drop gambar di sini</p>
                <p class="text-xs text-gray-400 mt-1">Bisa lebih dari satu gambar. PNG, JPG, JPEG (maks 5MB/file)</p>
              </div>
            </div>

            <!-- Publish Toggle -->
            <div class="flex items-center justify-between bg-gray-50 rounded-xl px-4 py-3">
              <div>
                <p class="text-sm font-semibold text-gray-700">Publikasikan Berita</p>
                <p class="text-xs text-gray-400 mt-0.5">Berita akan tampil di halaman publik</p>
              </div>
              <button
                type="button"
                @click="form.is_published = !form.is_published"
                :class="['relative w-12 h-6 rounded-full transition-colors duration-200',
                  form.is_published ? 'bg-green-500' : 'bg-gray-300']">
                <span :class="['absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full shadow transition-transform duration-200',
                  form.is_published ? 'translate-x-6' : 'translate-x-0']"></span>
              </button>
            </div>

            <!-- Error form -->
            <div v-if="formError" class="bg-red-50 text-red-600 text-sm px-4 py-3 rounded-xl border border-red-200">
              ❌ {{ formError }}
            </div>

            <!-- Actions -->
            <div class="flex gap-3 pt-2">
              <button type="button" @click="closeModal"
                class="flex-1 border border-gray-200 text-gray-600 py-2.5 rounded-xl text-sm font-semibold hover:bg-gray-50 transition">
                Batal
              </button>
              <button type="submit" :disabled="submitting"
                class="flex-1 bg-green-600 text-white py-2.5 rounded-xl text-sm font-semibold hover:bg-green-700 transition disabled:opacity-60 flex items-center justify-center gap-2">
                <div v-if="submitting" class="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                {{ submitting ? 'Menyimpan...' : (isEdit ? 'Simpan Perubahan' : 'Tambah Berita') }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Transition>

    <!-- ─── Modal Konfirmasi Hapus ──────────────────────────── -->
    <Transition name="modal">
      <div v-if="showDeleteModal"
        class="fixed inset-0 z-50 flex items-center justify-center p-4"
        @click.self="showDeleteModal = false">
        <div class="absolute inset-0 bg-black/50 backdrop-blur-sm"></div>
        <div class="relative bg-white rounded-2xl shadow-2xl w-full max-w-sm p-6 text-center">
          <p class="text-4xl mb-3">🗑️</p>
          <h3 class="text-lg font-black text-gray-900 mb-1">Hapus Berita?</h3>
          <p class="text-sm text-gray-500 mb-6">
            "<span class="font-semibold text-gray-700">{{ selectedNews?.title }}</span>"
            akan dihapus permanen.
          </p>
          <div class="flex gap-3">
            <button @click="showDeleteModal = false"
              class="flex-1 border border-gray-200 text-gray-600 py-2.5 rounded-xl text-sm font-semibold hover:bg-gray-50 transition">
              Batal
            </button>
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
import { ref, onMounted, nextTick } from 'vue'
import { useRouter, RouterLink } from 'vue-router'
import api from '../../services/api'

const BASE_URL = import.meta.env.VITE_BASE_URL || 'http://localhost:8080'
const router = useRouter()

// ─── State ───────────────────────────────────────────────────
const newsList   = ref([])
const loading    = ref(true)
const submitting = ref(false)
const showModal  = ref(false)
const showDeleteModal = ref(false)
const isEdit     = ref(false)
const selectedNews = ref(null)
const imageList  = ref([])
const contentEditor = ref(null)
const contentImageInput = ref(null)
const contentUploading = ref(false)
const formError  = ref('')
const alert      = ref({ show: false, type: 'success', message: '' })

const categories = ['Umum', 'Corporate News', 'Community Engagement', 'Sport & CSR', 'Sustainability']

const form = ref({
  title: '',
  slug: '',
  summary: '',
  content: '',
  category: '',
  thumbnail_path: '',
  published_at: '',
  is_published: false,
})

// ─── Helpers ─────────────────────────────────────────────────
const getImageUrl = (path) => {
  if (!path) {return ''}
  if (path.startsWith('http')) {return path}
  return `${BASE_URL}/${path.replace(/^\//, '')}`
}

const formatDate = (dateStr) => {
  if (!dateStr) {return '-'}
  return new Date(dateStr).toLocaleDateString('id-ID', {
    day: '2-digit', month: 'short', year: 'numeric'
  })
}

const formatDateInput = (dateStr) => {
  if (!dateStr) {return ''}
  return new Date(dateStr).toISOString().slice(0, 10)
}

const showAlert = (type, message) => {
  alert.value = { show: true, type, message }
  setTimeout(() => { alert.value.show = false }, 3000)
}

// ─── Fetch data ───────────────────────────────────────────────
const fetchNews = async () => {
  loading.value = true
  try {
    const res = await api.get('/admin/news')
    newsList.value = res.data?.data || res.data || []
  } catch (_e) {
    showAlert('error', 'Gagal memuat data berita')
  } finally {
    loading.value = false
  }
}

// ─── Modal Create ─────────────────────────────────────────────
const openCreate = async () => {
  isEdit.value = false
  form.value = { title: '', slug: '', summary: '', content: '', category: '', thumbnail_path: '', published_at: '', is_published: false }
  imageList.value = []
  formError.value = ''
  showModal.value = true
  await nextTick()
  syncEditorContent()
}

// ─── Modal Edit ───────────────────────────────────────────────
const openEdit = async (news) => {
  isEdit.value = true
  selectedNews.value = news
  form.value = {
    title: news.title || '',
    slug: news.slug || '',
    summary: news.summary || '',
    content: news.content || '',
    category: news.category || '',
    thumbnail_path: news.thumbnail_path || '',
    published_at: formatDateInput(news.published_at || news.PublishedAt),
    is_published: news.is_published || false,
  }
  imageList.value = (news.Images || []).map(img => ({ image_url: img.image_url }))
  if (!imageList.value.length && news.thumbnail_path) {
    imageList.value = [{ image_url: news.thumbnail_path }]
  }
  formError.value = ''
  showModal.value = true
  await nextTick()
  syncEditorContent()
}

const closeModal = () => {
  showModal.value = false
  formError.value = ''
}

// ─── Handle file upload ───────────────────────────────────────
const uploadSingleImage = async (file) => {
  const fd = new FormData()
  fd.append('image', file)
  const res = await api.post('/upload', fd, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
  return res.data?.url || ''
}

const uploadFiles = async (files) => {
  formError.value = ''
  for (const file of files) {
    if (!file.type.startsWith('image/')) {continue}
    if (file.size > 5 * 1024 * 1024) {
      formError.value = `File ${file.name} maksimal 5MB`
      continue
    }
    const url = await uploadSingleImage(file)
    imageList.value.push({ image_url: url })
  }
}

const handleFileChange = async (e) => {
  await uploadFiles(Array.from(e.target.files || []))
  e.target.value = ''
}

const handleDrop = async (e) => {
  await uploadFiles(Array.from(e.dataTransfer.files || []))
}

const removeImage = (index) => {
  imageList.value.splice(index, 1)
}

// ─── Rich text editor ──────────────────────────────────────
const execCmd = (cmd, val = null) => {
  document.execCommand(cmd, false, val)
  contentEditor.value?.focus()
}

const triggerContentImage = () => {
  contentImageInput.value?.click()
}

const insertContentImage = async (e) => {
  const file = e.target.files[0]
  if (!file) {return}
  if (file.size > 5 * 1024 * 1024) {
    formError.value = 'Ukuran gambar maksimal 5MB'
    return
  }
  contentUploading.value = true
  try {
    const url = await uploadSingleImage(file)
    document.execCommand('insertHTML', false, `<img src="${url}" class="max-w-full rounded-lg my-2" /><br/>`)
    onContentInput()
  } catch {
    formError.value = 'Gagal upload gambar'
  }
  contentUploading.value = false
  e.target.value = ''
}

const onContentInput = () => {
  form.value.content = contentEditor.value?.innerHTML || ''
}

const syncEditorContent = () => {
  if (contentEditor.value && form.value.content) {
    if (contentEditor.value.innerHTML !== form.value.content) {
      contentEditor.value.innerHTML = form.value.content
    }
  } else if (contentEditor.value) {
    contentEditor.value.innerHTML = form.value.content || ''
  }
}

// ─── Submit form (Create / Edit) ─────────────────────────────
const submitForm = async () => {
  if (!form.value.title.trim() || !form.value.content.trim()) {
    formError.value = 'Judul dan konten wajib diisi'
    return
  }
  submitting.value = true
  formError.value = ''
  try {
    form.value.thumbnail_path = imageList.value[0]?.image_url || ''

    // Auto-generate slug dari title jika kosong
    if (!form.value.slug) {
      form.value.slug = form.value.title
        .toLowerCase()
        .replace(/[^a-z0-9\s-]/g, '')
        .replace(/\s+/g, '-')
        .replace(/-+/g, '-')
        .trim()
    }

    const payload = {
      title:          form.value.title,
      slug:           form.value.slug,
      summary:        form.value.summary,
      content:        form.value.content,
      category:       form.value.category,
      thumbnail_path: form.value.thumbnail_path,
      published_at:   form.value.published_at,
      is_published:   form.value.is_published,
      Images:         imageList.value,
    }

    if (isEdit.value) {
      await api.put(`/admin/news/${selectedNews.value.ID}`, payload)
      showAlert('success', 'Berita berhasil diperbarui!')
    } else {
      await api.post('/admin/news', payload)
      showAlert('success', 'Berita berhasil ditambahkan!')
    }

    closeModal()
    await fetchNews()
  } catch (e) {
    formError.value = e.response?.data?.message || 'Terjadi kesalahan. Coba lagi.'
  } finally {
    submitting.value = false
  }
}

// ─── Delete ───────────────────────────────────────────────────
const confirmDelete = (news) => {
  selectedNews.value = news
  showDeleteModal.value = true
}

const doDelete = async () => {
  submitting.value = true
  try {
    await api.delete(`/admin/news/${selectedNews.value.ID}`)
    newsList.value = newsList.value.filter(n => n.ID !== selectedNews.value.ID)
    showDeleteModal.value = false
    showAlert('success', 'Berita berhasil dihapus!')
  } catch (_e) {
    showAlert('error', 'Gagal menghapus berita')
  } finally {
    submitting.value = false
  }
}

// ─── Logout ───────────────────────────────────────────────────
const logout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  router.push('/admin/login')
}

onMounted(() => {
  const token = localStorage.getItem('token')
  if (!token) { router.push('/admin/login'); return }
  fetchNews()
})
</script>

<style scoped>
.modal-enter-active, .modal-leave-active {
  transition: opacity 0.2s ease;
}
.modal-enter-from, .modal-leave-to {
  opacity: 0;
}
[contenteditable]:empty:before {
  content: attr(data-placeholder);
  color: #9ca3af;
  pointer-events: none;
}
</style>