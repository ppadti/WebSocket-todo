import {
  Button,
  Card,
  Flex,
  Form,
  TextContent,
  Text,
  TextVariants,
  TextInput,
  FlexItem,
} from "@patternfly/react-core";
import React, { FormEvent, useState } from "react";
import { Todo } from "./Model";
import { CheckIcon } from "@patternfly/react-icons";

type Props = {
  todo: Todo;
  handleEdit: (id: number, todo: string) => void;
  handleDone: (id: number, isDone: boolean, todo: string) => void;
  handleDelete: (id: number, todo: string) => void;
};

const SingleTodo = ({ todo, handleEdit, handleDone, handleDelete }: Props) => {
  if (!todo) {
    return null;
  }
  const [edit, setEdit] = useState<boolean>(false);
  const [editTodo, setEditTodo] = useState<string>(todo.todo);

  const handleEditOption = (e: FormEvent, id: number) => {
    e.preventDefault();
    handleEdit(id, editTodo);
    setEdit(false);
  };

  const updateTask = (name: string) => {
    setEditTodo(name);
  };

  return (
    <>
      {/* {!todo && <p>hello</p>} */}
      <Card style={{ width: "20rem", padding: "1rem", margin: "5rem" }}>
        <Form
          style={{ padding: "1rem" }}
          onSubmit={(e) => {
            console.log(todo.id);
            handleEditOption(e, todo.id);
          }}
        >
          {edit ? (
            <Flex
              display={{ default: "inlineFlex" }}
              flexWrap={{ default: "wrap" }}
            >
              <FlexItem>
                <TextInput
                  value={editTodo}
                  type="text"
                  id="horizontal-form-name"
                  aria-describedby="horizontal-form-name-helper"
                  name="horizontal-form-name"
                  onChange={updateTask}
                />
              </FlexItem>
              <FlexItem>
                <CheckIcon
                  onClick={(e) => {
                    handleEditOption(e, todo.id);
                  }}
                />
              </FlexItem>
            </Flex>
          ) : todo.isDone ? (
            <s>{todo.todo}</s>
          ) : (
            <TextContent>
              <Text component={TextVariants.h2}>{todo.todo}</Text>
            </TextContent>
          )}
        </Form>
        <Flex>
          <Button
            className="option"
            onClick={() => {
              if (!todo.isDone) setEdit(!edit);
            }}
          >
            Edit
          </Button>
          <Button
            variant="secondary"
            isDanger
            onClick={() => {
              handleDelete(todo.id, todo.todo);
            }}
          >
            Delete
          </Button>
          <Button
            variant="primary"
            onClick={() => {
              handleDone(todo.id, todo.isDone, todo.todo);
            }}
          >
            Done
          </Button>
        </Flex>
      </Card>
    </>
  );
};

export default SingleTodo;
