<!-- src/views/Home.vue -->

<template>
  <div>
    <h1>カレンダー</h1>
    <Calendar /> 
  </div>
</template>

<script>
import Calendar from '@/components/Calendar'; 

export default {
  components: {
    Calendar,
  },
};
</script>