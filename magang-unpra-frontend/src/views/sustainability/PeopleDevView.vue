<template>

  <!-- ── HERO ── -->
  <section class="relative h-[420px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('https://images.unsplash.com/photo-1522071820081-009f0129c71c?w=1920')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/40 to-transparent"></div>

    <div class="absolute top-10 left-10 flex items-center gap-2">
      <span class="w-1 h-5 bg-green-400 rounded-full"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>

    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4">
        People <span class="text-green-400">Development</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed">
        Preparing the best talent to achieve the highest performance and business sustainability.
      </p>
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400">
        <span class="hover:text-white cursor-pointer transition">Home</span>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">People Development</span>
      </div>
    </div>
  </section>

  <!-- ── STAT BAR ── -->
  <section class="bg-green-700 py-5">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-3 divide-x divide-green-600">
        <div class="text-center px-6">
          <p class="text-2xl font-black text-white">5 <span class="text-green-300 text-base font-semibold">Programs</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">Development Pillars</p>
        </div>
        <div class="text-center px-6">
          <p class="text-2xl font-black text-white">100% <span class="text-green-300 text-base font-semibold">Committed</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">To Excellence</p>
        </div>
        <div class="text-center px-6">
          <p class="text-2xl font-black text-white">Skills <span class="text-green-300 text-base font-semibold">+ Values</span></p>
          <p class="text-green-200 text-xs tracking-widest uppercase mt-1">Integrated Approach</p>
        </div>
      </div>
    </div>
  </section>

  <!-- ── PHOTO GRID ── -->
  <section class="py-16 bg-white">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-3 gap-3 h-64 rounded-2xl overflow-hidden">
        <div class="col-span-1 overflow-hidden">
          <img src="https://images.unsplash.com/photo-1573496359142-b8d87734a5a2?w=600"
            alt="Training" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
        <div class="col-span-1 overflow-hidden">
          <img src="https://images.unsplash.com/photo-1522071820081-009f0129c71c?w=600"
            alt="Workshop" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
        <div class="col-span-1 overflow-hidden">
          <img src="https://images.unsplash.com/photo-1517048676732-d65bc937f952?w=600"
            alt="Team" class="w-full h-full object-cover hover:scale-105 transition-transform duration-500"/>
        </div>
      </div>
    </div>
  </section>

  <!-- ── INTRO TEXT ── -->
  <section class="pb-16 bg-white">
    <div class="max-w-5xl mx-auto px-10">
      <div class="grid grid-cols-5 gap-12">

        <!-- Teks utama -->
        <div class="col-span-3 space-y-4 text-gray-500 text-sm leading-relaxed">
          <p>
            People development is carried out in order to prepare the best talent to be able to show the
            best performance and to achieve the business targets they carry. Talking about business is not
            only about running but also how they manage the business even more and sustain it in the future.
          </p>
          <p>
            PT Tanjungenim Lestari Pulp and Paper are committed to developing competent, motivated and
            integrated information systems to achieve a competitive organization. The commitment is realized
            in several activities and achievements as follows:
          </p>
        </div>

        <!-- Highlight box -->
        <div class="col-span-2 bg-green-50 border border-green-100 rounded-2xl p-6 flex flex-col justify-center">
          <div class="w-10 h-10 rounded-xl bg-green-100 flex items-center justify-center mb-4">
            <span class="text-green-600 text-lg">🌱</span>
          </div>
          <p class="text-green-800 font-bold text-base leading-snug mb-2">
            Competent. Motivated. Integrated.
          </p>
          <p class="text-green-600 text-xs leading-relaxed">
            Building people who are ready to sustain and grow the business for the long term.
          </p>
        </div>

      </div>
    </div>
  </section>

  <!-- ── COMMITMENT LIST ── -->
  <section class="py-16 bg-gray-50">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-12">
        <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3">
          Our Commitments
        </span>
        <h2 class="text-3xl font-black text-gray-900">Development Pillars</h2>
        <p class="text-gray-400 text-sm mt-2 max-w-md mx-auto">
          Five core programs that shape TEL's approach to people development.
        </p>
      </div>

      <div class="space-y-4">
        <div v-for="(item, index) in commitments" :key="index"
          class="flex gap-5 bg-white border border-gray-100 rounded-2xl p-6 hover:shadow-md hover:border-green-100 transition group">

          <!-- Nomor -->
          <div class="w-10 h-10 rounded-xl bg-green-600 text-white font-black text-sm flex items-center justify-center flex-shrink-0 group-hover:bg-green-700 transition">
            {{ String(index + 1).padStart(2, '0') }}
          </div>

          <!-- Teks -->
          <div>
            <h3 class="text-gray-900 font-bold text-sm mb-1">{{ item.title }}</h3>
            <p class="text-gray-400 text-sm leading-relaxed">{{ item.desc }}</p>
          </div>
        </div>
      </div>

    </div>
  </section>

  <!-- ── IMAGE SLIDER ── -->
  <section class="py-20 bg-white">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-10">
        <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3">
          Gallery
        </span>
        <h2 class="text-3xl font-black text-gray-900">People in Action</h2>
      </div>

      <!-- Slider wrapper -->
      <div class="relative rounded-2xl overflow-hidden shadow-xl aspect-video bg-gray-100">

        <!-- Gambar aktif -->
        <transition name="fade" mode="out-in">
          <img
            :key="activeSlide"
            :src="slides[activeSlide].url"
            :alt="slides[activeSlide].caption"
            class="w-full h-full object-cover"
          />
        </transition>

        <!-- Caption -->
        <div class="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-6">
          <p class="text-white text-sm font-semibold">{{ slides[activeSlide].caption }}</p>
          <p class="text-gray-300 text-xs mt-1">{{ activeSlide + 1 }} / {{ slides.length }}</p>
        </div>

        <!-- Tombol Prev -->
        <button @click="prev"
          class="absolute left-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-black/50 text-white flex items-center justify-center hover:bg-black/70 transition backdrop-blur-sm">
          ‹
        </button>

        <!-- Tombol Next -->
        <button @click="next"
          class="absolute right-4 top-1/2 -translate-y-1/2 w-10 h-10 rounded-full bg-black/50 text-white flex items-center justify-center hover:bg-black/70 transition backdrop-blur-sm">
          ›
        </button>
      </div>

      <!-- Dot Pagination -->
      <div class="flex justify-center gap-2 mt-6">
        <button
          v-for="(_, i) in slides" :key="i"
          @click="activeSlide = i"
          :class="[
            'transition-all duration-300 rounded-full',
            activeSlide === i
              ? 'bg-green-600 w-6 h-2.5'
              : 'bg-gray-200 hover:bg-gray-300 w-2.5 h-2.5'
          ]">
        </button>
      </div>

    </div>
  </section>

  <!-- ── CTA ── -->
  <section class="py-16 bg-green-700">
    <div class="max-w-5xl mx-auto px-10 text-center">
      <h2 class="text-3xl font-black text-white mb-3">
        Growing People, Growing Business
      </h2>
      <p class="text-green-200 text-sm mb-8 max-w-lg mx-auto">
        TEL continuously invests in its people to ensure sustainable business excellence and community growth.
      </p>
      
    </div>
  </section>

</template>

<script setup>
import { ref } from 'vue'

const activeSlide = ref(0)

const commitments = [
  {
    title: 'Skills-Based Leadership Coaching',
    desc: 'Achieve a high performing organization by developing skills-based Coaching for leadership styles through a comprehensive Supervisory / Management / Leadership Program.'
  },
  {
    title: 'Succession Planning & Career Path',
    desc: 'Develop and implement succession plans for critical positions to ensure the regeneration of employees through the Career Path and Talent Management.'
  },
  {
    title: 'Community & Scholarship Programs',
    desc: 'Prepare a generation of quality sources from internal employees as well as from the village community through scholarships, internship, community education, school foundation programs and infrastructure improvement.'
  },
  {
    title: 'Technology & Information Skills',
    desc: 'Improve knowledge and skills related to technology and information across all levels of the organization.'
  },
  {
    title: 'Rules, Competencies & Knowledge',
    desc: 'Meet the rules and competencies with good knowledge and skills to maintain professional standards throughout the company.'
  },
]

const slides = [
  {
    url: 'https://images.unsplash.com/photo-1573496359142-b8d87734a5a2?w=1200',
    caption: 'Leadership Training Session'
  },
  {
    url: 'https://images.unsplash.com/photo-1522071820081-009f0129c71c?w=1200',
    caption: 'Team Collaboration Workshop'
  },
  {
    url: 'https://images.unsplash.com/photo-1517048676732-d65bc937f952?w=1200',
    caption: 'Management Development Program'
  },
  {
    url: 'https://images.unsplash.com/photo-1531482615713-2afd69097998?w=1200',
    caption: 'Skills Training & Certification'
  },
  {
    url: 'https://images.unsplash.com/photo-1543269865-cbf427effbad?w=1200',
    caption: 'Community Education Initiative'
  },
]

function next() {
  activeSlide.value = (activeSlide.value + 1) % slides.length
}
function prev() {
  activeSlide.value = (activeSlide.value - 1 + slides.length) % slides.length
}
</script>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.4s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>