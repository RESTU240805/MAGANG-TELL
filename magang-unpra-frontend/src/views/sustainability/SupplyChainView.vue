<template>

  <!-- ── HERO ── -->
  <section class="relative h-[380px] flex items-end bg-cover bg-center overflow-hidden"
    style="background-image: url('https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?w=1920')">
    <div class="absolute inset-0 bg-gradient-to-t from-black/85 via-black/45 to-transparent"></div>

    <div class="absolute top-10 left-10 flex items-center gap-2">
      <span class="w-1 h-5 bg-green-400 rounded-full"></span>
      <span class="text-green-400 text-xs font-bold tracking-[3px] uppercase">Sustainability</span>
    </div>

    <div class="relative z-10 max-w-5xl mx-auto px-10 pb-14 w-full">
      <h1 class="text-5xl font-black text-white leading-tight mb-4">
        Supply Chain <span class="text-green-400">Management</span>
      </h1>
      <p class="text-gray-300 text-sm max-w-xl leading-relaxed">
        A strategic, sustainable, and responsible approach to managing our entire supply chain operations.
      </p>
      <div class="flex items-center gap-2 mt-5 text-xs text-gray-400">
        <span class="hover:text-white cursor-pointer transition">Home</span>
        <span class="text-gray-600">/</span>
        <span class="hover:text-white cursor-pointer transition">Sustainability</span>
        <span class="text-gray-600">/</span>
        <span class="text-green-400 font-semibold">Supply Chain Management</span>
      </div>
    </div>
  </section>

  

  <!-- ── ACCORDION SECTION ── -->
  <section class="py-20 bg-gray-50">
    <div class="max-w-5xl mx-auto px-10">

      <div class="text-center mb-12">
        <span class="inline-block bg-green-100 text-green-700 text-xs font-bold tracking-widest uppercase px-3 py-1 rounded-full mb-3">
          Our Framework
        </span>
        <h2 class="text-3xl font-black text-gray-900">Supply Chain Pillars</h2>
        <p class="text-gray-400 text-sm mt-2 max-w-md mx-auto">
          Click each section to explore our complete supply chain management framework.
        </p>
      </div>

      <div class="space-y-4">

        <!-- ── ACCORDION 1: Strategy ── -->
        <div class="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden">
          <button
            @click="toggle(0)"
            class="w-full flex items-center justify-between px-7 py-5 text-left transition hover:bg-gray-50"
            :class="open[0] ? 'bg-green-600 hover:bg-green-700' : 'bg-white'">
            <div class="flex items-center gap-4">
              
              <span :class="open[0] ? 'text-white' : 'text-gray-900'"
                class="font-bold text-base transition">
                Strategy and Strong Management Plan
              </span>
            </div>
            <svg :class="[open[0] ? 'rotate-180 text-white' : 'text-gray-400', 'transition-transform duration-300 w-5 h-5 flex-shrink-0']"
              fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>

          <transition name="accordion">
            <div v-if="open[0]" class="px-7 pb-7 pt-2">
              <p class="text-gray-500 text-sm leading-relaxed mb-5">
                We have a strategy for strong management plan as follows:
              </p>

              <!-- Card: Stable Pulp Production -->
              <div class="bg-green-50 border border-green-100 rounded-xl p-5 mb-4">
                <div class="flex items-start gap-3 mb-3">
                  <span class="w-2 h-2 rounded-full bg-green-500 mt-1.5 flex-shrink-0"></span>
                  <p class="text-gray-700 text-sm font-semibold">
                    To keep stable pulp production and its good quality which consist of:
                  </p>
                </div>
                <div class="grid grid-cols-2 gap-2 ml-5">
                  <div v-for="item in strategyItems" :key="item"
                    class="flex items-center gap-2 bg-white rounded-lg px-4 py-2.5 border border-green-100 shadow-sm">
                    <span class="w-1.5 h-1.5 rounded-full bg-green-400 flex-shrink-0"></span>
                    <span class="text-sm text-gray-600">{{ item }}</span>
                  </div>
                </div>
              </div>

              <!-- Item 2 & 3 -->
              <div class="space-y-3">
                <div class="flex items-start gap-3 bg-gray-50 rounded-xl px-5 py-3.5 border border-gray-100">
                  <span class="w-2 h-2 rounded-full bg-green-500 mt-1.5 flex-shrink-0"></span>
                  <p class="text-gray-600 text-sm">
                    To implement sound Industrial Relation, Corporate Social Responsibility.
                  </p>
                </div>
                <div class="flex items-start gap-3 bg-gray-50 rounded-xl px-5 py-3.5 border border-gray-100">
                  <span class="w-2 h-2 rounded-full bg-green-500 mt-1.5 flex-shrink-0"></span>
                  <p class="text-gray-600 text-sm">
                    To develop competent and motivated People.
                  </p>
                </div>
              </div>
            </div>
          </transition>
        </div>

        <!-- ── ACCORDION 2: Sustainability ── -->
        <div class="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden">
          <button
            @click="toggle(1)"
            class="w-full flex items-center justify-between px-7 py-5 text-left transition hover:bg-gray-50"
            :class="open[1] ? 'bg-green-600 hover:bg-green-700' : 'bg-white'">
            <div class="flex items-center gap-4">
              
              <span :class="open[1] ? 'text-white' : 'text-gray-900'"
                class="font-bold text-base transition">
                Supply Chain Sustainability
              </span>
            </div>
            <svg :class="[open[1] ? 'rotate-180 text-white' : 'text-gray-400', 'transition-transform duration-300 w-5 h-5 flex-shrink-0']"
              fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>

          <transition name="accordion">
            <div v-if="open[1]" class="px-7 pb-7 pt-2">

              <!-- Motto highlight -->
              <div class="bg-green-700 rounded-xl p-5 mb-5 text-center">
                <p class="text-green-200 text-xs font-bold tracking-widest uppercase mb-1">TEL Motto</p>
                <p class="text-white text-lg font-black">"Fairness, Innovation & Harmony"</p>
              </div>

              <div class="space-y-4">
                <div v-for="(item, i) in sustainabilityItems" :key="i"
                  class="flex gap-4 bg-gray-50 rounded-xl p-5 border border-gray-100 hover:border-green-100 hover:bg-green-50/50 transition">
                  <div class="w-8 h-8 rounded-full bg-green-600 text-white font-black text-xs flex items-center justify-center flex-shrink-0 mt-0.5">
                    {{ i + 1 }}
                  </div>
                  <p class="text-gray-600 text-sm leading-relaxed">{{ item }}</p>
                </div>
              </div>
            </div>
          </transition>
        </div>

        <!-- ── ACCORDION 3: Policy ── -->
        <div class="bg-white rounded-2xl border border-gray-100 shadow-sm overflow-hidden">
          <button
            @click="toggle(2)"
            class="w-full flex items-center justify-between px-7 py-5 text-left transition hover:bg-gray-50"
            :class="open[2] ? 'bg-green-600 hover:bg-green-700' : 'bg-white'">
            <div class="flex items-center gap-4">
              
              <span :class="open[2] ? 'text-white' : 'text-gray-900'"
                class="font-bold text-base transition">
                Policy & Guidelines
              </span>
            </div>
            <svg :class="[open[2] ? 'rotate-180 text-white' : 'text-gray-400', 'transition-transform duration-300 w-5 h-5 flex-shrink-0']"
              fill="none" stroke="currentColor" stroke-width="2.5" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7"/>
            </svg>
          </button>

          <transition name="accordion">
            <div v-if="open[2]" class="px-7 pb-7 pt-4">
              <div class="space-y-5">
                <div v-for="(policy, i) in policies" :key="i"
                  class="border border-gray-100 rounded-xl overflow-hidden">

                  <!-- Policy header -->
                  <div class="flex items-center gap-3 bg-gray-50 px-5 py-3 border-b border-gray-100">
                    <span class="w-7 h-7 rounded-lg bg-green-600 text-white font-black text-xs flex items-center justify-center flex-shrink-0">
                      {{ i + 1 }}
                    </span>
                    <h4 class="font-bold text-gray-900 text-sm">{{ policy.title }}</h4>
                  </div>

                  <!-- Policy items -->
                  <div class="px-5 py-4 space-y-2">
                    <div v-for="(point, j) in policy.points" :key="j"
                      class="flex items-start gap-2.5">
                      <span class="w-1.5 h-1.5 rounded-full bg-green-400 mt-1.5 flex-shrink-0"></span>
                      <p class="text-gray-500 text-sm leading-relaxed">{{ point }}</p>
                    </div>

                    <!-- Sub-section untuk Information Disclosure -->
                    <div v-if="policy.procedures" class="mt-4 ml-4 border-l-2 border-green-200 pl-4 space-y-3">
                      <p class="text-gray-600 text-xs font-semibold uppercase tracking-widest mb-2">
                        Procedures for vendors not meeting labor standards:
                      </p>
                      <div v-for="(proc, k) in policy.procedures" :key="k"
                        class="flex gap-3 bg-green-50 rounded-lg p-3">
                        <span class="w-5 h-5 rounded-full bg-green-600 text-white font-bold text-[10px] flex items-center justify-center flex-shrink-0 mt-0.5">
                          {{ k + 1 }}
                        </span>
                        <p class="text-gray-600 text-xs leading-relaxed">{{ proc }}</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </transition>
        </div>

      </div>
    </div>
  </section>

  <!-- ── CTA ── -->
  <section class="py-16 bg-green-700">
    <div class="max-w-5xl mx-auto px-10 text-center">
      <h2 class="text-3xl font-black text-white mb-3">
        Committed to Responsible Supply Chain
      </h2>
      <p class="text-green-200 text-sm mb-8 max-w-lg mx-auto">
        TEL continuously upholds fairness, innovation, and harmony across every link in its supply chain.
      </p>
    
    </div>
  </section>

</template>

<script setup>
import { ref } from 'vue'

const open = ref([true, false, false])

function toggle(index) {
  open.value[index] = !open.value[index]
}

const strategyItems = [
  'Log Management',
  'Procurement Management',
  'Spare Part Management',
  'Contract Management',
]

const sustainabilityItems = [
  'PT.Tanjungenim Lestari Pulp and Paper has a principle stated in a motto "Fairness, Innovation & Harmony" which underlines all of its activities and transactions that occur to ensure the sustainability of business activities in harmony.',
  'PT.Tanjungenim Lestari Pulp and Paper commits to uphold Good Corporate Governance (GCG) principles in its operation. To achieve this, the participation of all stakeholders including suppliers is very essential. In the implementation stage, PT.Tanjungenim Lestari Pulp and Paper appreciate any party who provide feedback or information when contrary event to the principle of GCG are found, such as: request of offer of bribery, extortion to suppliers, etc.',
  'PT.Tanjungenim Lestari Pulp and Paper is part of Marubeni Group who already set its supply chain sustainability guidelines. PT.Tanjungenim Lestari Pulp and Paper supports strengthening sustainability throughout its supply chain with the objective of building an environmentally friendly, healthy and sustainable society.',
  'Cascading the Marubeni Group supply chain sustainability guidelines, PT.Tanjungenim Lestari Pulp and Paper asks for the understanding and cooperation of its business partners in observing the Guidelines, to promote highly effective sustainability initiatives and assist business partners\' capacity building by making recommendations and sharing examples of good practices.',
]

const policies = [
  {
    title: 'Observance of Laws',
    points: [
      'Observe the laws of the countries where business is conducted and laws relating to business transactions.',
    ],
  },
  {
    title: 'Respect for Human Rights',
    points: [
      'Respect human rights without discrimination, harassment of any kind, abuse or other inhumane treatment.',
      'No child labor or forced labor.',
      'Proper management of employees\' work hours, breaks and holidays and prohibition of excessive overtime work.',
      'Payment of the legally mandated minimum wage and endeavoring to pay at least a living wage. No inappropriate wage abatement.',
      'Respect for employees\' right to unionize for the purpose of labor-management negotiations and right to collective bargaining.',
    ],
  },
  {
    title: 'Conservation of the Environment',
    points: [
      'Recognize that climate change issues are important and respond appropriately.',
      'Protect the natural environment.',
      'Reduce environmental negative impact; prevent pollution.',
    ],
  },
  {
    title: 'Fair Transactions',
    points: [
      'Conduct fair transactions and do not inhibit free competition.',
      'Prevent corruption; offer no bribes or illegal contributions.',
    ],
  },
  {
    title: 'Safety and Health',
    points: [
      'Ensure safe and healthy workplaces and maintain a good working environment.',
    ],
  },
  {
    title: 'Quality Control',
    points: [
      'Maintain the quality and safety of products and services.',
    ],
  },
  {
    title: 'Information Disclosure',
    points: [
      'Timely and appropriate disclosure of information.',
      'As part of this Policy, the Marubeni Group has set out the following procedures for dealing with vendors that do not meet labor standards (Observance of Laws, Respect for Human Rights, Safety and Health):',
    ],
    procedures: [
      'When it comes to light that a vendor has failed to meet labor standards, we will as necessary ask the vendor to ascertain the facts, and if the facts are true, prepare a report on the background of the issue and improvement measures. Depending on circumstances, we may also inspect the vendor.',
      'If we determine that improvement measures are insufficient, we will request that further measures be taken.',
      'If, despite implementing steps (i) and (ii) above, the situation does not improve, we will examine whether to continue our relationship with the vendor.',
    ],
  },
]
</script>

<style scoped>
.accordion-enter-active,
.accordion-leave-active {
  transition: all 0.35s ease;
  overflow: hidden;
}
.accordion-enter-from,
.accordion-leave-to {
  opacity: 0;
  max-height: 0;
}
.accordion-enter-to,
.accordion-leave-from {
  opacity: 1;
  max-height: 2000px;
}
</style>