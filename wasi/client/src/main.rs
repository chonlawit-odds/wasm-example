use wasmtime::*;
use wasmtime_wasi::sync::WasiCtxBuilder;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // สร้าง Engine และสร้าง WASI context
    let engine = Engine::default();
    let wasi_ctx = WasiCtxBuilder::new()
        .inherit_stdout() // ใช้ stdout ของระบบ
        .inherit_stderr() // ใช้ stderr ของระบบ
        .build();

    // สร้าง Store พร้อม WASI context
    let mut store = Store::new(&engine, wasi_ctx);

    // โหลดโมดูล WebAssembly จากไฟล์ main.wasm
    let module = Module::from_file(&engine, "main.wasm")?;

    // สร้าง Linker และเพิ่ม WASI environment
    let mut linker = Linker::new(&engine);
    wasmtime_wasi::add_to_linker(&mut linker, |ctx| ctx)?;

    // สร้าง instance จากโมดูล
    let instance = linker.instantiate(&mut store, &module)?;

    // เรียกฟังก์ชัน add จาก instance
    let add = instance.get_typed_func::<(i32, i32), i32>(&mut store, "add")?;

    // ทดสอบเรียกฟังก์ชัน add
    let result = add.call(&mut store, (2025, 543))?;
    println!("Result of add(2025, 543): {}", result);

    Ok(())
}
