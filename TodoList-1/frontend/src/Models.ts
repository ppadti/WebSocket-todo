export interface Todo {
  id: number;
  todo: string;
  isDone: boolean;
}
export type setTodo = (todo: Todo) => void;
