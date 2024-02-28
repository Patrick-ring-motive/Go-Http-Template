package submodules




func Try(try func(),catch func(interface{})){
  defer func() {
      if e := recover(); e != nil {
          catch(e)
      }
  }()
  try()
}

func Catch(catch func(interface{}))func(interface{}){
  return catch
}




func IfTry(ifbool func()bool,try func(),elsecatch func(interface{})){
  defer func() {
      if e := recover(); e != nil {
          elsecatch(e)
      }
  }()
  if(ifbool()){
    try()
  }else{
    elsecatch(nil)
  }
}

func ElseCatch(catch func(interface{}))func(interface{}){
  return catch
}
