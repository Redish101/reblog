import Button from "@/components/Button";
import TextField from "@/components/TextField";

export default function App() {
  return (
    <div className="flex gap-1 p-24">
      <div>
        <h2 className="font-light text-3xl">Buttons</h2>
        <div className="flex gap-1">
          <Button variant="filled">Button</Button>
          <Button variant="outlined">Button</Button>
          <Button variant="text">Button</Button>
        </div>
        <h2 className="font-light text-3xl mt-12">Text Fields</h2>
        <div className="flex gap-1">
          <TextField
            label="Label"
            placeholder="Placeholder"
            variant="filled"
          />
        </div>
      </div>
    </div>
  );
}
