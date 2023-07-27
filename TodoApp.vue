<template>
  <div style="max-width: 600px">
    <h2>My Vue Todo App</h2>

    <div>
      <input
        type="text"
        v-model="task"
        placeholder="Enter task"
      />
      <button @click="submitTask()">
        SUBMIT
      </button>
    </div>

    <table class="table table-bordered mt-5">
      <thead>
        <tr>
          <th scope="col">Task</th>
          <th scope="col" style="width: 120px">Status</th>
          <th scope="col" class="text-center">Delete</th>
          <th scope="col" class="text-center">Edit</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(task, index) in tasks" :key="index">
          <td>
            <span :class="{ 'line-through': task.status === 'finished' }">
              {{ task.name }}
            </span>
          </td>
          <td>
            <span
              class="pointer noselect"
              @click="update(index)"
              :class="{
                'text1': task.status === 'pre-started',
                'text2': task.status === 'finished',
                'text3': task.status === 'pending',
              }"
            >
              {{ capitalize(task.status) }}
            </span>
          </td>
          <td class="text-center">
            <div @click="del(index)">
              <button>Delete</button>
            </div>
          </td>
          <td class="text-center">
            <div @click="edit(index)">
              <button>Edit</button>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: "HelloWorld",

  data() {
    return {
      task: "",
      editedTask: null,
      statuses: ["pre-started", "pending", "finished"],

      tasks: [
        {
          name: "Example 1",
          status: "pre-started",
        },
        {
          name: "Example 2",
          status: "pending",
        },
        {
          name: "Example 3",
          status: "finished",
        },
      ],
    };
  },

  methods: {

    capitalize(str) {
      return str.charAt(0).toUpperCase() + str.slice(1);
    },

    update(index) {
      let newIndex = this.statuses.indexOf(this.tasks[index].status);
      if (++newIndex > 2) newIndex = 0;
      this.tasks[index].status = this.statuses[newIndex];
    },

    del(index) {
      this.tasks.splice(index, 1);
    },

    edit(index) {
      this.task = this.tasks[index].name;
      this.editedTask = index;
    },

    submit() {
      if (this.task.length === 0) return;
      if (this.editedTask != null) {
        this.tasks[this.editedTask].name = this.task;
        this.editedTask = null;
      } else {
        this.tasks.push({
          name: this.task,
          status: "pre-started",
        });
      }

      this.task = "";
    },
  },
};
</script>

<style scoped>
.pointer {
  cursor: pointer;
}
.line-through {
  text-decoration: line-through;
}
</style>