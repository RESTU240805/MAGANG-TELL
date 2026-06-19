<template>
  <div class="rich-editor border border-gray-200 rounded-xl overflow-hidden focus-within:ring-2 focus-within:ring-green-500 transition-shadow">
    <!-- Toolbar -->
    <div class="flex items-center gap-1 px-3 py-2 bg-gray-50 border-b border-gray-200 flex-wrap">
      <button
        type="button"
        @click="editor?.chain().focus().toggleBold().run()"
        :class="['p-1.5 rounded-lg text-sm font-bold transition cursor-pointer', editor?.isActive('bold') ? 'bg-green-100 text-green-700' : 'text-gray-600 hover:bg-gray-200']"
        title="Bold"
      >B</button>

      <button
        type="button"
        @click="editor?.chain().focus().toggleItalic().run()"
        :class="['p-1.5 rounded-lg text-sm italic transition cursor-pointer', editor?.isActive('italic') ? 'bg-green-100 text-green-700' : 'text-gray-600 hover:bg-gray-200']"
        title="Italic"
      >I</button>

      <div class="w-px h-5 bg-gray-300 mx-1"></div>

      <button
        type="button"
        @click="editor?.chain().focus().toggleBulletList().run()"
        :class="['p-1.5 rounded-lg transition cursor-pointer', editor?.isActive('bulletList') ? 'bg-green-100 text-green-700' : 'text-gray-600 hover:bg-gray-200']"
        title="Bullet List"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/></svg>
      </button>

      <button
        type="button"
        @click="editor?.chain().focus().toggleOrderedList().run()"
        :class="['p-1.5 rounded-lg transition cursor-pointer', editor?.isActive('orderedList') ? 'bg-green-100 text-green-700' : 'text-gray-600 hover:bg-gray-200']"
        title="Ordered List"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6h13M7 12h13M7 18h13M3 6h.01M3 12h.01M3 18h.01"/></svg>
      </button>

      <div class="w-px h-5 bg-gray-300 mx-1"></div>

      <button
        type="button"
        @click="editor?.chain().focus().toggleCodeBlock().run()"
        :class="['p-1.5 rounded-lg font-mono text-xs transition cursor-pointer', editor?.isActive('codeBlock') ? 'bg-green-100 text-green-700' : 'text-gray-600 hover:bg-gray-200']"
        title="Code Block"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/></svg>
      </button>

      <div class="w-px h-5 bg-gray-300 mx-1"></div>

      <button
        type="button"
        @click="triggerImageUpload"
        class="p-1.5 rounded-lg text-gray-600 hover:bg-gray-200 transition flex items-center gap-1.5 cursor-pointer"
        title="Insert Image"
      >
        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>
        <span class="text-xs font-medium">Gambar</span>
      </button>
      <input ref="imageInput" type="file" accept="image/*" class="hidden" @change="handleImageUpload" />
    </div>

    <!-- Editor Content -->
    <editor-content :editor="editor" class="prose-editor min-h-[200px] max-h-[500px] overflow-y-auto px-4 py-3 text-sm text-gray-700 focus:outline-none" />
  </div>
</template>

<script setup>
import { ref, watch, onBeforeUnmount } from 'vue'
import { useEditor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Image from '@tiptap/extension-image'
import Placeholder from '@tiptap/extension-placeholder'
import api from '../services/api'

const props = defineProps({
  modelValue: { type: String, default: '' },
  placeholder: { type: String, default: 'Tulis deskripsi di sini...' }
})

const emit = defineEmits(['update:modelValue'])
const imageInput = ref(null)

const editor = useEditor({
  content: props.modelValue,
  extensions: [
    StarterKit,
    Image.configure({ inline: false, allowBase64: true }),
    Placeholder.configure({ placeholder: props.placeholder })
  ],
  onUpdate: ({ editor }) => {
    emit('update:modelValue', editor.getHTML())
  }
})

watch(() => props.modelValue, (val) => {
  if (editor.value && val !== editor.value.getHTML()) {
    editor.value.commands.setContent(val, false)
  }
})

const triggerImageUpload = () => {
  imageInput.value?.click()
}

const handleImageUpload = async (e) => {
  const file = e.target.files?.[0]
  if (!file) return

  if (file.size > 5 * 1024 * 1024) {
    alert('Ukuran gambar maksimal 5MB')
    return
  }

  try {
    const formData = new FormData()
    formData.append('image', file)
    const res = await api.post('/upload', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
    const url = res.data.url
    if (url) {
      editor.value?.chain().focus().setImage({ src: url }).run()
    }
  } catch {
    alert('Gagal upload gambar')
  }

  if (imageInput.value) imageInput.value.value = ''
}

onBeforeUnmount(() => {
  editor.value?.destroy()
})
</script>

<style>
.prose-editor .tiptap {
  outline: none;
  min-height: 200px;
}
.prose-editor .tiptap p.is-editor-empty:first-child::before {
  color: #9ca3af;
  content: attr(data-placeholder);
  float: left;
  height: 0;
  pointer-events: none;
}
.prose-editor .tiptap h1 { font-size: 1.5em; font-weight: 700; margin: 0.5em 0; }
.prose-editor .tiptap h2 { font-size: 1.25em; font-weight: 600; margin: 0.5em 0; }
.prose-editor .tiptap h3 { font-size: 1.1em; font-weight: 600; margin: 0.5em 0; }
.prose-editor .tiptap ul { list-style-type: disc; padding-left: 1.5em; margin: 0.5em 0; }
.prose-editor .tiptap ol { list-style-type: decimal; padding-left: 1.5em; margin: 0.5em 0; }
.prose-editor .tiptap li { margin: 0.25em 0; }
.prose-editor .tiptap pre {
  background: #1e293b;
  color: #e2e8f0;
  padding: 12px 16px;
  border-radius: 8px;
  font-family: monospace;
  font-size: 0.875em;
  overflow-x: auto;
  margin: 0.75em 0;
}
.prose-editor .tiptap code {
  background: #f1f5f9;
  padding: 2px 6px;
  border-radius: 4px;
  font-family: monospace;
  font-size: 0.9em;
}
.prose-editor .tiptap pre code {
  background: transparent;
  padding: 0;
}
.prose-editor .tiptap img {
  max-width: 100%;
  border-radius: 8px;
  margin: 1em 0;
}
.prose-editor .tiptap blockquote {
  border-left: 3px solid #10b981;
  padding-left: 12px;
  margin: 0.75em 0;
  color: #6b7280;
}
</style>
