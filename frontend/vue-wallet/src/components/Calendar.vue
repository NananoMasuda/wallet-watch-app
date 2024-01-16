<!-- src/components/Calendar.vue -->

<template>
  <div>
    <h2>{{ currentYear }} {{ currentMonth }}</h2>
    <button @click="prevMonth">前の月</button>
    <button @click="nextMonth">次の月</button>
    <div class="calendar" v-for="(week, index) in calendars" :key="index">
        <div v-for="(day, index) in week" :key="index">
            <p>{{ day.date }}</p>
        </div>
    </div>
  </div>
</template>

<script>
import moment from "moment";
export default {
  data() {
    return {
      currentDate: moment(),
      currentMonth: moment().format("MMM"),
      currentYear: moment().format("YYYY"),
    };
  },
  mounted(){
    console.log(this.getCalendar());
  },
  computed: {
    calendars() {
      return this.getCalendar();
    },
  },
  methods: {
    getStartDate() {
    
      let date = moment(this.currentDate); // 今日の日付を取得
      date.startOf("month"); // 月の最初の日付を取得、dateを上書き
      const weekdayNum = date.day(); // 月の最初の曜日ナンバーを取得（0が日曜、1が月曜、2が火曜...）

      return  moment(date).subtract(weekdayNum, "days"); // 最初の日付より過去の日曜日

    },
    getEndDate() {

      let date = moment(this.currentDate); // 今日の日付を取得
      date.endOf("month"); // 月の最後の日付を取得
      const weekdayNum = date.day(); // 月の最後の曜日ナンバーを取得

      return  moment(date).add(6 - weekdayNum, "days"); // 最後の日付より未来の日曜日
    },
    getCalendar() {

      let startDate = this.getStartDate();
      const endDate = this.getEndDate();
      const weekNumber = Math.ceil(endDate.diff(startDate, "days") / 7); // カレンダーの高さ

      let calendars = [];

      for (let week = 0; week < weekNumber; week++) {
        let weekRow = [];
        for (let day = 0;  day < 7; day++) {
            weekRow.push({
            date: startDate.get("date"),
            });
            startDate.add(1, "days");
        }
        calendars.push(weekRow);
      }

      return calendars;

    },
    setCurrentMonthAndYear() {

      this.currentMonth = moment(this.currentDate).format("MMM");
      this.currentYear = moment(this.currentDate).format("YYYY");

    },
    nextMonth() {

      this.currentDate = moment(this.currentDate).add(1, "month");
      this.setCurrentMonthAndYear();

    },
    prevMonth() {

      this.currentDate = moment(this.currentDate).subtract(1, "month");
      this.setCurrentMonthAndYear();

    },
  }
}
</script>

<style>
  .calendar{
    display:flex;
  }

  .calendar div {
    width: 220px;
    height: 150px;
    border: 1px solid #ccc;
  }

  .calendar div p{
    color: #85878b;
    margin: 4px 0px 0px 0px;
  }

  button {
    margin-bottom: 15px;
    margin-left: 6px;
    width: 100px;
    height: 50px;
    font-size: 14px;
  }
</style>