import {
  Box,
  Button,
  ListItem,
  ListItemSecondaryAction,
  ListItemText,
  Paper,
  TextField,
} from "@mui/material";
import React, { useState } from "react";
import { Todo } from "../Models";

type Props = {
  todo: Todo;
  handleUpdate: (id: number, name: string) => void;
  handleDelete: (id: number) => void;
};

const SingleTodo = ({ todo, handleUpdate, handleDelete }: Props) => {
  const [editedName, setEditedName] = useState<string>(todo.todo);
  const [edit, setEdit] = useState(false);

  const handleCancel = () => {
    setEdit(false);
  };

  const handleSave = (taskId: number) => {
    console.log(taskId);
    setEdit(false);
    handleUpdate(taskId, editedName);
  };

  return (
    <Paper
      sx={{
        border: "1px",
        borderRadius: "5px",
        marginBottom: "15px",
        padding: "5px",
      }}
    >
      <ListItem>
        {edit ? (
          <Box key={todo.id}>
            <TextField
              sx={{ width: "330px" }}
              label="Todo"
              value={editedName}
              onChange={(e) => {
                setEditedName(e.target.value);
              }}
            />
            <Button
              variant="contained"
              sx={{ padding: "5px", margin: "10px" }}
              onClick={() => {
                handleSave(todo.id);
              }}
            >
              Save
            </Button>
            <Button variant="contained" onClick={handleCancel}>
              Cancel
            </Button>
          </Box>
        ) : (
          <>
            <ListItemText primary={todo.todo} />
            <ListItemSecondaryAction>
              <Button
                variant="contained"
                onClick={() => {
                  setEdit(true);
                }}
                sx={{ margin: "5px" }}
                size="small"
              >
                Edit
              </Button>
              <Button
                variant="contained"
                onClick={() => {
                  handleDelete(todo.id);
                }}
                size="small"
              >
                Delete
              </Button>
            </ListItemSecondaryAction>
          </>
        )}
      </ListItem>
    </Paper>
  );
};

export default SingleTodo;
