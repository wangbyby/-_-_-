from __future__ import print_function
import keras
from keras.models import Sequential, load_model
from keras.layers import Dense, Dropout, Flatten
from keras.layers import Conv2D, MaxPooling2D
from keras.preprocessing import  image
import  numpy as np
from PIL import Image
import os
batch_size = 128
num_classes = 10
epochs = 12

img_rows, img_cols = 28, 28
classs = {1:'狗',0:'猫'}
"""
train文件夹下
    0 : 猫🐱
    1 : 狗🐕
"""

#裁剪图片
def convert_one_image(file_path,width=100,height=100):
    img = Image.open(file_path)
    try:
        new_img = img.resize((width,height), Image.BILINEAR)
        new_img.save(file_path)
    except Exception as e:
        print("convert image error")
        print(e)
def convert_all_images(all_images_path):
    for i in os.listdir(all_images_path):
        _file_path = all_images_path + '/' + i
        for j in os.listdir(_file_path):

            path_image = _file_path + '/' + j
            convert_one_image(path_image)  # 转换
def images_nparray(all_images_path):
    import os
    _x = []
    _y = []
    for i in os.listdir(all_images_path):
        _file_path = all_images_path+'/'+i
        for j in os.listdir(_file_path):
            y_label = j.split('_')[0]
            _y.append(y_label)
            path_image = _file_path +'/'+ j

            #图片到np数组
            img__ = Image.open(path_image).convert('RGB')
            img__ = np.array(img__)
            _x.append(img__)
    return np.array(_x), np.array(_y)
def nparray_to_npz(train_path, test_path,npz_file_path):
    convert_all_images(train_path)
    convert_all_images(test_path)
    x_train,y_train = images_nparray(train_path)
    x_test,y_test = images_nparray(test_path)
    #保存压缩文件
    np.savez(npz_file_path, x_train=x_train,y_train=y_train,x_test=x_test,y_test=y_test)



#训练函数
def running_train_test(npz_file_path,save_file_path):
    input_shape = (100,100,3)
    # the data, split between train and test sets
    data = np.load(npz_file_path)
    x_train, y_train, x_test, y_test = data['x_train'],data['y_train'],data['x_test'],data['y_test']

    x_train = x_train.astype('float32')
    x_test = x_test.astype('float32')
    x_train /= 255
    x_test /= 255
    print('x_train shape:', x_train.shape)
    print(x_train.shape[0], 'train samples')
    print(x_test.shape[0], 'test samples')

    # convert class vectors to binary class matrices
    y_train = keras.utils.to_categorical(y_train, num_classes)
    y_test = keras.utils.to_categorical(y_test, num_classes)

    model = Sequential()
    model.add(Conv2D(32, kernel_size=(3, 3),
                     activation='relu',
                     input_shape=input_shape))
    model.add(Conv2D(64, (3, 3), activation='relu'))
    model.add(MaxPooling2D(pool_size=(2, 2)))
    model.add(Dropout(0.25))
    model.add(Flatten())
    model.add(Dense(128, activation='relu'))
    model.add(Dropout(0.5))
    model.add(Dense(num_classes, activation='softmax'))

    model.compile(loss=keras.losses.categorical_crossentropy,
                  optimizer=keras.optimizers.Adadelta(),
                  metrics=['accuracy'])

    model.fit(x_train, y_train,
              batch_size=batch_size,
              epochs=epochs,
              verbose=1,
              validation_data=(x_test, y_test))
    score = model.evaluate(x_test, y_test, verbose=0)
    print('Test loss:', score[0])
    print('Test accuracy:', score[1])
    model.save(save_file_path)

def train_test_images_to_model(train, test, npz_path, h5_path):

    nparray_to_npz(train,test,npz_path)
    running_train_test(npz_path,h5_path)

def image_file_nparray(image_path, width=100,height=100):
    img = image.load_img(image_path, target_size=(width,height))
    img = image.img_to_array(img)
    x = np.expand_dims(img, axis=0)
    return x

def load_model_from_h5_and_predict(h5_path, pre_path):
    model = load_model(h5_path)
    convert_one_image(pre_path)
    pre_nparray = image_file_nparray(pre_path)
    pre_y = model.predict_classes(pre_nparray)
    print("结果 : ",pre_y)
    print("预测结果 : ",classs[pre_y[0]])

# def predict_result_to_cat_or_dog(pre_y):
#     d= {1:0,0:0}
#     for i in pre_y:
#         d[i] +=1
#     sum_c = len(pre_y)
#     cat = d[0]/sum_c
#     dog = d[1]/sum_c
#     print("狗的概率 : ",dog)
#     print("猫的概率 : ",cat)
 #main函数...
if "__main__" == __name__:
    npz_path = 'cat_dog.npz'
    save_path = 'save.h5'
    test_dir_path = 'test'
    train_dir_path = 'train'
    train_test_images_to_model(train_dir_path, test_dir_path, npz_path, save_path)
    pre_path = 'pre.jpg'
    # load_model_from_h5_and_predict(save_path, pre_path)
