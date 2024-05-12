import Button from "./components/Button";

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
      </div>
    </div>
  );
}
