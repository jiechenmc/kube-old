"use client"
import Image from "next/image";
import React, { FC, useState } from "react";
import { ReactSortable } from "react-sortablejs";

interface ItemType {
  id: number;
  name: string;
}
export default function Home() {
  const [state, setState] = useState([
    { id: 1, name: "shrek" },
    { id: 2, name: "fiona" },
  ]);

  const [state2, setState2] = useState([
    { id: 3, name: "shrek" },
    { id: 4, name: "fiona" },
  ]);

  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <ReactSortable list={state} setList={setState} animation={200} delay={2} group="shared-group-name"  >
        {state.map((item) => (
          <div key={item.id}>{item.name}</div>
        ))}
      </ReactSortable>
      <ReactSortable list={state2} setList={setState2} animation={200} delay={2} group="shared-group-name"  >
        {state2.map((item) => (
          <div key={item.id}>{item.name}</div>
        ))}
      </ReactSortable>
    </div>
  );
}
