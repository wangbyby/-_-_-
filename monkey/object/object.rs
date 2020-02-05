use crate::ast::ast;
use crate::evaluator::evaluator::Environment;
type ObjectType= &'static str;

pub const INTEGER_OBJ:ObjectType = "INTEGER";
pub const BOOLEAN_OBJ: ObjectType = "BOOLEAN";
pub const NULL_OBJ: ObjectType = "NULL";
pub const RETURN_VALUE_OBJ:ObjectType = "RETURN_VALUE";
pub const ERROR_OBJ: ObjectType = "ERROR";
pub const FUNCTION_OBJ: ObjectType = "FUNCTION";


#[derive(Debug,Clone,PartialEq)]
pub enum TheObject {
    Integer(i64),
    Boolean(bool),
    ReturnValue(Box<TheObject>),
    Errors(String),
    Func(Vec<Option<Box<ast::Identifier>>>, Box<ast::ASTNode>,  Environment),
    NULL,

}

impl TheObject {
    pub fn default()->Self{
        TheObject::NULL
    }
    pub fn type_of(&self)->ObjectType{
        use self::TheObject::*;
        match self {
            Integer(_)=> INTEGER_OBJ,
            Boolean(_)=> BOOLEAN_OBJ,
            ReturnValue(_)=> RETURN_VALUE_OBJ,
            Errors(_)=> ERROR_OBJ,
            Func(_,_,_)=>FUNCTION_OBJ,
            NULL => NULL_OBJ,
        }
    }

    pub fn inspect(&self)->String {
        use self::TheObject::*;
        match self{
            Integer(i)=> format!("{}",*i),
            Boolean(i)=> format!("{}",*i),
            ReturnValue(i)=> format!("{}",i.as_ref().inspect()),
            Errors(i)=> i.clone(),
            Func(ref ident, ref block,ref env) => format!("env = {:?} fn({:?}){}\n{:?}\n{} ",env,ident, "{",block,"}"),
            NULL => format!("()"),
        }
    }

    pub fn is_error(&self)->bool{
        self.type_of() == ERROR_OBJ
    }
}



// use std::collections::HashMap;
// pub struct Environment{
//     store: HashMap<String,TheObject>,
// }

// impl Environment{
//     pub fn default()->Self{
//         Environment{
//             store: HashMap::new(),
//         }
//     }

//     pub fn get<S: Into<String>>(&self,k:S, v: TheObject){
//         self.store.in
//     }
// }











// pub trait Object:std::fmt::Debug {
//     fn type_of(&self) ->ObjectType;
//     fn inspect(&self) ->String;
//     fn as_any(&self) ->&dyn Any;
// }
// #[derive(Debug,Copy,Clone)]
// pub struct Integer{
//     pub value:i64,
// }

// impl Object for Integer{
//     fn inspect(&self) ->String{
//         format!("{}",self.value)
//     }
//     fn type_of(&self) ->ObjectType{
//         INTEGER_OBJ
//     }
//     fn as_any(&self)->&dyn Any{
//         self
//     }
// }

// #[derive(Debug,Copy,Clone)]
// pub struct Boolean{
//     pub value:bool,
// }
// impl Object for Boolean{
//     fn inspect(&self) ->String{
//         format!("{}",self.value)
//     }
//     fn type_of(&self) ->ObjectType{
//         BOOLEAN_OBJ
//     }
//     fn as_any(&self)->&dyn Any{
//         self
//     }
// }

// #[derive(Debug,Copy,Clone)]
// pub struct Null{}

// impl Object for Null{
//     fn inspect(&self) ->String{
//         "null".to_string()
//     }
//     fn type_of(&self) ->ObjectType{
//         NULL_OBJ
//     }
//     fn as_any(&self)->&dyn Any{
//         self
//     }
// }

