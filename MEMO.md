entity.Content みたいなものを作って、Post と Content を分離すると Repository の実装が分離しない。  
しかし、Entity を正しく設計した結果としてそうなるのであって、Repository の実装を分離したくないから Entity を細かくするという順序になるのは間違い（Entity が先、Repository が後）
