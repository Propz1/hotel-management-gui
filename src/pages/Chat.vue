<template>
  <div :class="$style.chat"><SidebarDark5 /></div>


  <div id="calendarTable" >
		<h1>My Calendar</h1>

 

      <div :class="$style.calendarparent">
    
		<calendar-view
      :items="items"
			:show-date="showDate"
      :time-format-options="{ hour: 'numeric', minute: '2-digit' }"
			:enable-drag-drop="true"
			:disable-past="disablePast"
			:disable-future="disableFuture"
			:show-times="showTimes"
			:date-classes="myDateClasses"
			:display-period-uom="displayPeriodUom"
			:display-period-count="displayPeriodCount"
			:starting-day-of-week="startingDayOfWeek"
			:class="themeClasses"
			:period-changed-callback="periodChanged"
			:current-period-label="useTodayIcons ? 'icons' : ''"
			:displayWeekNumbers="displayWeekNumbers"
			:enable-date-selection="true"
			:selection-start="selectionStart"
			:selection-end="selectionEnd"
			@date-selection-start="setSelection"
			@date-selection="setSelection"
			@date-selection-finish="finishSelection"
			@drop-on-date="onDrop"
			@click-date="onClickDay"
			@click-item="onClickItem"
			class="theme-default holiday-us-traditional holiday-us-official">

			<template #header="{ headerProps }">
				<calendar-view-header
					:header-props="headerProps"
					@input="setShowDate"> 
        </calendar-view-header>
			</template>

		</calendar-view>
  </div>
	</div>

  





</template>
<script>
  import { defineComponent } from "vue";
  import SidebarDark5 from "../components/SidebarDark5.vue";
  import { CalendarView, CalendarViewHeader,CalendarMath} from "vue-simple-calendar";
  import '../defaultCalendar.css';
  import "vue-simple-calendar/dist/css/default.css"
  import "vue-simple-calendar/dist/css/holidays-us.css"

  export default defineComponent({

    data() {
        return {
          showDate: this.thisMonth(1),
          message: "",
			    startingDayOfWeek: 0,
			    disablePast: false,
			    disableFuture: false,
			    displayPeriodUom: "month",
			    displayPeriodCount: 1,
			    displayWeekNumbers: false,
			    showTimes: true,
			    selectionStart: null,
			    selectionEnd: null,
			    newItemTitle: "",
			    newItemStartDate: "",
			    newItemEndDate: "",
			    useDefaultTheme: true,
			    useHolidayTheme: true,
			    useTodayIcons: false,
          items: [
				/*{
					id: "e0",
					startDate: "2018-01-05",
				},*/
				{
					id: "e1",
					startDate: this.thisMonth(15, 18, 30),
				},
				{
					id: "e2",
					startDate: this.thisMonth(15),
					title: "Single-day item with a long title",
				},
				{
					id: "e3",
					startDate: this.thisMonth(7, 9, 25),
					endDate: this.thisMonth(10, 16, 30),
					title: "Multi-day item with a long title and times",
				},
				{
					id: "e4",
					startDate: this.thisMonth(20),
					title: "My Birthday!",
					classes: "birthday",
					url: "https://en.wikipedia.org/wiki/Birthday",
				},
				{
					id: "e5",
					startDate: this.thisMonth(5),
					endDate: this.thisMonth(12),
					title: "Multi-day item",
					classes: "purple",
					tooltip: "This spans multiple days",
				},
				{
					id: "foo",
					startDate: this.thisMonth(29),
					title: "Same day 1",
				},
				{
					id: "e6",
					startDate: this.thisMonth(29),
					title: "Same day 2",
					classes: "orange",
				},
				{
					id: "e7",
					startDate: this.thisMonth(29),
					title: "Same day 3",
				},
				{
					id: "e8",
					startDate: this.thisMonth(29),
					title: "Same day 4",
					classes: "orange",
				},
				{
					id: "e9",
					startDate: this.thisMonth(29),
					title: "Same day 5",
				},
				{
					id: "e10",
					startDate: this.thisMonth(29),
					title: "Same day 6",
					classes: "orange",
				},
				{
					id: "e11",
					startDate: this.thisMonth(29),
					title: "Same day 7",
				},
			],

        };

    },

    name: "Chat",

    components: {
       SidebarDark5,
       CalendarView,
			 CalendarViewHeader, 
       CalendarMath, 
      },

       
    computed: {
		   userLocale() {
		    	return CalendarMath.getDefaultBrowserLocale
		   },

		   dayNames() {
			    return CalendarMath.getFormattedWeekdayNames(this.userLocale, "long", 0)
		   },

		   themeClasses() {
			     return {
				   "theme-default": this.useDefaultTheme,
				   "holiday-us-traditional": this.useHolidayTheme,
				   "holiday-us-official": this.useHolidayTheme,
			    }
		   },

		   myDateClasses() {
			   // This was added to demonstrate the dateClasses prop. Note in particular that the
			   // keys of the object are `yyyy-mm-dd` ISO date strings (not dates), and the values
			   // for those keys are strings or string arrays. Keep in mind that your CSS to style these
			   // may need to be fairly specific to make it override your theme's styles. See the
			   // CSS at the bottom of this component to see how these are styled.
		   	 const o = {}
		     const theFirst = this.thisMonth(1)
			   const ides = [2, 4, 6, 9].includes(theFirst.getMonth()) ? 15 : 13
			   const idesDate = this.thisMonth(ides)
		     o[CalendarMath.isoYearMonthDay(idesDate)] = "ides"
		   	 o[CalendarMath.isoYearMonthDay(this.thisMonth(21))] = ["do-you-remember", "the-21st"]
			   return o
	    	},
	 },

	    mounted() {
	    	this.newItemStartDate = CalendarMath.isoYearMonthDay(CalendarMath.today())
		    this.newItemEndDate = CalendarMath.isoYearMonthDay(CalendarMath.today())
	    },


    methods: {
      onLinkMainPageContainerClick() {
        this.$router.push("/main");
      },
      onLinkCalendarPageContainerClick() {
        this.$router.push("/calendar");
      },
      onLinkRequisitionsPageContainerClick() {
        this.$router.push("/");
      },
      onLinkOccupancyPageContainerClick() {
        this.$router.push("/requisitions");
      },
      onLinkChatPageContainerClick() {
        this.$router.push("/chat");
      },
      onLinkReportsPageContainerClick() {
        this.$router.push("/reports");
      },
      onLinkIntegrationPageContainerClick() {
        this.$router.push("/integrations");
      },
      onLinkHelpPageContainerClick() {
        this.$router.push("/help");
      },
      onLinkSettingsPageContainerClick() {
        this.$router.push("/settings");
      },


      setShowDate(d) {
				this.showDate = d;
			},

      periodChanged() {
			// range, eventSource) {
			// Demo does nothing with this information, just including the method to demonstrate how
			// you can listen for changes to the displayed range and react to them (by loading items, etc.)
			//console.log(eventSource)
			//console.log(range)
		  },

		  thisMonth(d, h, m) {
			   const t = new Date()
			   return new Date(t.getFullYear(), t.getMonth(), d, h || 0, m || 0)
		   },

		  onClickDay(d) {
		   	this.selectionStart = null
		   	this.selectionEnd = null
			   this.message = `You clicked: ${d.toLocaleDateString()}`
		  },

		  onClickItem(e) {
		  	this.message = `You clicked: ${e.title}`
		  },

		  setShowDate(d) {
		   	this.message = `Changing calendar view to ${d.toLocaleDateString()}`
		   	this.showDate = d
		  },

		  setSelection(dateRange) {
			   this.selectionEnd = dateRange[1]
			   this.selectionStart = dateRange[0]
		  },

	    finishSelection(dateRange) {
		   	this.setSelection(dateRange)
			  this.message = `You selected: ${this.selectionStart.toLocaleDateString()} -${this.selectionEnd.toLocaleDateString()}`
		  },

		  onDrop(item, date) {
		   	this.message = `You dropped ${item.id} on ${date.toLocaleDateString()}`
		   	// Determine the delta between the old start date and the date chosen,
			   // and apply that delta to both the start and end date to move the item.
		   	const eLength = CalendarMath.dayDiff(item.startDate, date)
		   	item.originalItem.startDate = CalendarMath.addDays(item.startDate, eLength)
			   item.originalItem.endDate = CalendarMath.addDays(item.endDate, eLength)
		  },

		  clickTestAddItem() {
		    	this.items.push({
				  startDate: this.newItemStartDate,
				  endDate: this.newItemEndDate,
				  title: this.newItemTitle,
				  id: "e" + Math.random().toString(36).substr(2, 10),
			})

			this.message = "You added a calendar item!"
		},

    },
  });
</script>

<style module>
  .chat {
    position: relative;
    background-color: var(--color-whitesmoke-300);
    width: 100%;
    height: 1080px;
    display: flex;
    flex-direction: row;
    align-items: flex-start;
    justify-content: flex-start;
  }


  #calendarTable {
		/* font-family: 'Avenir', Helvetica, Arial, sans-serif;
		color: #2c3e50;
		height: 67vh;
		width: 90vw;
		margin-left: auto;
		margin-right: auto; */



    /* position: absolute; */
    /* top: 68px;
    left: 249px; */
    /* width: 89.5%;
    height: 1010px;  */

    /* display: flex;
    flex-direction: column;
    flex-grow: 1; */

    display: flex;
	flex-direction: row;
	font-family: Calibri, sans-serif;
	width: 95vw;
	min-width: 30rem;
	max-width: 100rem;
	min-height: 40rem;
	margin-left: auto;
	margin-right: auto;
   
	}

  .calendarparent {
	display: flex;
	flex-direction: column;
	flex-grow: 1;
	overflow-x: hidden;
	overflow-y: hidden;
	max-height: 80vh;
	background-color: white;
}

  .cv-header-days {
    /* display: flex;
    flex-grow: 1;
    flex-shrink: 0;
    flex-basis: 0;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
    text-align: center;
    border-width: 1px 1px 0 0; */

  /* display: flex;
	flex-grow: 0;
	flex-shrink: 0;
	flex-basis: auto;
	flex-flow: row nowrap;
	border-width: 0 0 0 1px; */

  /* display: flex;
    flex-grow: 1;
    flex-shrink: 0;
    flex-basis: 0;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
    text-align: center;
    border-width: 1px 1px 0 0; */

    display: flex;
    flex-grow: 0;
    flex-shrink: 0;
    flex-basis: auto;
    flex-flow: row nowrap;
    border-width: 0 0 0 1px;

    flex-direction: column;
    flex-grow: 1;
}

.cvwrapper{
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    height: 100%;
    min-height: 100%;
    max-height: 100%;
    overflow-x: hidden;
    overflow-y: hidden
}


</style>
